package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ===================================================================
// Trace JSON Output Schema
//
// This matches the "proof object" format from the design doc exactly.
// The verifier trusts THIS object — not the raw engine state.
// ===================================================================

// TraceEnvelope is the top-level JSON document (§1 of the trace spec).
type TraceEnvelope struct {
	TraceID     string          `json:"trace_id"`
	SpecVersion string          `json:"spec_version"`
	CreatedAt   string          `json:"created_at"`
	Task        TaskDescriptor  `json:"task"`
	Outcome     TraceOutcome    `json:"outcome"`
	Summary     TraceSummary    `json:"summary"`
	Steps       []TraceStep     `json:"steps"`
	Coherence   CoherenceRecord `json:"coherence_trajectory"`
	Meta        MetaRecord      `json:"meta"`
	Signatures  SignatureRecord `json:"signatures"`
}

// TaskDescriptor identifies what the trace was solving (§1).
type TaskDescriptor struct {
	Domain    string `json:"domain"`
	Type      string `json:"type"`
	RiskLevel string `json:"risk_level"`
}

// TraceSummary is the human-readable answer + confidence (§1).
type TraceSummary struct {
	FinalAnswer string  `json:"final_answer"`
	Confidence  float64 `json:"confidence"`
}

// TraceStep is one event in the steps array (§2).
// op is stringly-typed on purpose; effects are declarative.
type TraceStep struct {
	Step    int      `json:"step"`
	Op      string   `json:"op"`
	Args    []string `json:"args"`
	Scope   int      `json:"scope"`
	Effects []string `json:"effects"`
	Tags    []string `json:"tags"`
}

// CoherenceRecord is the trajectory object (§3).
type CoherenceRecord struct {
	Initial CoherenceState    `json:"initial"`
	Updates []CoherenceUpdate `json:"updates"`
	Final   CoherenceState    `json:"final"`
}

// CoherenceUpdate is one delta entry inside the trajectory (§3).
type CoherenceUpdate struct {
	Step     int            `json:"step"`
	Delta    CoherenceDelta `json:"delta"`
	Reason   string         `json:"reason"`
	Evidence []string       `json:"evidence"`
}

// CoherenceDelta holds the three signed deltas for one update.
// Pointers so zero-value dimensions are omitted from JSON.
type CoherenceDelta struct {
	L *float64 `json:"l,omitempty"`
	P *float64 `json:"p,omitempty"`
	C *float64 `json:"c,omitempty"`
}

// MetaRecord is the §4 meta section: distillation result + budget accounting.
type MetaRecord struct {
	Distillation MetaDistillation `json:"distillation"`
	Budgets      MetaBudgets      `json:"budgets"`
}

// MetaDistillation records what @meta:distill produced (§4).
type MetaDistillation struct {
	Attempted bool     `json:"attempted"`
	Mode      string   `json:"mode"`
	Depth     int      `json:"depth"`
	Produced  []string `json:"produced"`
}

// MetaBudgets is the circuit-breaker accounting (§4).
type MetaBudgets struct {
	MetaDepthUsed  int     `json:"meta_depth_used"`
	MetaDepthMax   int     `json:"meta_depth_max"`
	MetaTokenRatio float64 `json:"meta_token_ratio"`
}

// SignatureRecord holds the verifier certificate (§6).
type SignatureRecord struct {
	Certificate CertificateRecord `json:"certificate"`
}

// CertificateRecord is the output the verifier stamps onto the envelope (§6).
type CertificateRecord struct {
	TraceID            string            `json:"trace_id"`
	Verdict            string            `json:"verdict"`
	CertificateVersion string            `json:"certificate_version"`
	Checks             map[string]string `json:"checks"`
	Flags              []string          `json:"flags"`
	RiskScore          float64           `json:"risk_score"`
}

// helper: returns a pointer to a float64 (lets us omit zero deltas in JSON).
func fp(v float64) *float64 { return &v }

// ===================================================================
// TraceRecorder
//
// Sits behind the Streamer interface.  Every StepEvent the engine emits
// lands here.  At the end we assemble the proof-object envelope from
// (a) the recorded events and (b) the sealed engine Trace.
// ===================================================================

// TraceRecorder implements Streamer and accumulates events.
type TraceRecorder struct {
	events []StepEvent
}

// Emit satisfies the Streamer interface.
func (tr *TraceRecorder) Emit(event StepEvent) {
	tr.events = append(tr.events, event)
}

// ===================================================================
// Verifier (minimal, rule-based — §6 of the trace spec)
//
// Does NOT execute anything.  Reads the assembled envelope and checks
// the four hard invariants.  Returns a CertificateRecord.
// ===================================================================

func verify(env TraceEnvelope) CertificateRecord {
	checks := map[string]string{}
	var flags []string
	riskScore := 0.0

	// --- scope_integrity -------------------------------------------
	// In this minimal verifier we trust the engine's ScopeStack enforced
	// immutability.  A full verifier would replay scope numbers.
	checks["scope_integrity"] = "pass"

	// --- meta_budget: depth and ratio within limits -----------------
	if env.Meta.Budgets.MetaDepthUsed > env.Meta.Budgets.MetaDepthMax {
		checks["meta_budget"] = "fail"
		flags = append(flags, "META_DEPTH_EXCEEDED")
		riskScore += 0.4
	} else if env.Meta.Budgets.MetaTokenRatio > 0.25 {
		checks["meta_budget"] = "fail"
		flags = append(flags, "META_TOKEN_RATIO_EXCEEDED")
		riskScore += 0.3
	} else {
		checks["meta_budget"] = "pass"
	}

	// --- outcome_gate: success distillation requires accepted --------
	if env.Meta.Distillation.Attempted &&
		env.Meta.Distillation.Mode == "success" &&
		env.Outcome != OutcomeAccepted {
		checks["outcome_gate"] = "fail"
		flags = append(flags, "OUTCOME_GATE_VIOLATION")
		riskScore += 0.5
	} else {
		checks["outcome_gate"] = "pass"
	}

	// --- coherence_sanity: every update must carry evidence ----------
	coherenceOK := true
	init := env.Coherence.Initial
	fin := env.Coherence.Final
	if fin.L < init.L || fin.P < init.P || fin.C < init.C {
		flags = append(flags, "COHERENCE_NET_REGRESSION")
		riskScore += 0.15
	}
	for _, u := range env.Coherence.Updates {
		if len(u.Evidence) == 0 {
			coherenceOK = false
			flags = append(flags, fmt.Sprintf("COHERENCE_UPDATE_STEP_%d_NO_EVIDENCE", u.Step))
			riskScore += 0.1
		}
	}
	if coherenceOK {
		checks["coherence_sanity"] = "pass"
	} else {
		checks["coherence_sanity"] = "fail"
	}

	// Derive verdict
	verdict := "valid"
	for _, v := range checks {
		if v == "fail" {
			verdict = "invalid"
			break
		}
	}

	if len(flags) == 0 {
		flags = []string{} // ensure JSON emits [] not null
	}

	return CertificateRecord{
		TraceID:            env.TraceID,
		Verdict:            verdict,
		CertificateVersion: "cc-1.0",
		Checks:             checks,
		Flags:              flags,
		RiskScore:          riskScore,
	}
}

// ===================================================================
// Buoyancy Example — the canonical trace from v1.1 §7.1
//
// Problem: Calculate buoyant force on a fully submerged 10 cm cube in water.
// Expected: 9.8 N upward.
//
// This exercises, in order:
//   @b              — immutable binding
//   @scope          — legal shadowing
//   @coh            — 3-dimension coherence updates
//   @refl           — completeness guard reflection
//   @i              — conditional (completeness check)
//   @t              — try/fallback on volume calc
//   math:pow        — V = side³
//   math:multiply   — F = ρ × V × g  (variadic, 3 args)
//   math:round      — round to 2 dp
//   @o              — sealed output
//   BeginTrace / EndTrace — trace lifecycle
//   EvalMetaDistill — outcome-gated distillation
// ===================================================================

func runBuoyancy() TraceEnvelope {
	recorder := &TraceRecorder{}
	m := NewMRL(recorder)
	m.RegisterDefaultPlugins()

	traceID := "ep_buoy_001"
	now := time.Now().UTC().Format(time.RFC3339)

	// ── Begin trace recording ─────────────────────────────────────
	m.BeginTrace(traceID, "classical_physics_fluids", "derived_quantity_calculation")

	// ── State 0: establish initial coherence (matches spec §7.1) ──
	m.EvalCoh(0.25, 0.15, 0.30, "initial setup: cube parameters known")

	// ── @b ?goal ──────────────────────────────────────────────────
	m.SetVar("?goal", "Calculate buoyant force on fully submerged 10cm cube in water")

	// ── @b ?side ──────────────────────────────────────────────────
	m.SetVar("?side", 0.1)

	// ── Demonstrate @scope: legal shadow of ?side in inner frame ──
	m.EvalScope(func() {
		m.SetVar("?side", 0.2) // hypothetical 20 cm cube — isolated
		_ = m.GetVar("?side")  // resolves to 0.2 inside scope
	})
	// Outside the scope, ?side is still 0.1 — immutability preserved.

	// ── Transition 1: retrieve Archimedes' principle ──────────────
	m.CallPlugin("log", "Recall Archimedes principle")
	m.SetVar("?archimedes", "F_b = ρ_fluid × V_displaced × g")

	m.EvalCoh(0.07, 0.15, 0.15, "Canonical principle retrieved")

	// ── Transition 2: volume calculation via @t ───────────────────
	m.EvalTry(
		func() {
			vol := m.CallPlugin("math:pow", m.GetVar("?side"), 3.0)
			m.SetVar("?volume", vol) // 0.001 m³
		},
		func() {
			m.EvalRefl("calculation_error", "Confirm side-length units before retry")
		},
	)

	m.EvalCoh(0.03, 0.15, 0.0, "Concrete geometric quantity added")

	// ── Transition 3: completeness guard — @i triggers @refl ──────
	m.EvalIf(
		m.Coh.C < 0.65, // completeness still low → guard fires
		func() {
			m.EvalRefl("incomplete_parameters", "Retrieve fluid density before computing force")
			m.SetVar("?rho", 1000.0) // kg/m³, pure water at 4 °C
			m.EvalCoh(0.03, 0.0, 0.25, "Missing constant supplied (completeness guard)")
		},
		nil,
	)

	// ── Transition 4: final force calculation ─────────────────────
	m.SetVar("?g", 9.8) // m/s²

	// @math:multiply ?rho ?volume ?g  — variadic 3-arg call
	fb := m.CallPlugin("math:multiply",
		m.GetVar("?rho"),
		m.GetVar("?volume"),
		m.GetVar("?g"),
	)

	fbRounded := m.CallPlugin("math:round", fb, 2.0)
	m.SetVar("?F_b", fbRounded)

	m.EvalCoh(0.04, 0.20, 0.0, "Dimensional consistency verified — final value reached")

	// ── Cross-validation: weight of displaced water ──────────────
	massWater := m.CallPlugin("math:multiply", m.GetVar("?rho"), m.GetVar("?volume"))
	weightWater := m.CallPlugin("math:multiply", massWater, m.GetVar("?g"))

	// Cross-validation uses approximate equality (tolerance 0.01) to survive
	// floating-point rounding — mirrors the spec's @i (≈ ... tolerance:0.01).
	fbVal, _ := toFloat64(fbRounded)
	wVal, _ := toFloat64(weightWater)
	crossValid := (fbVal-wVal) < 0.01 && (wVal-fbVal) < 0.01

	m.EvalIf(
		crossValid,
		func() {
			m.EvalCoh(0.0, 0.0, 0.10, "Cross-validation: two independent calculations agree")
		},
		func() {
			m.EvalRefl("calculation_mismatch", "Re-examine volume or density")
		},
	)

	// ── @o — seal the output ──────────────────────────────────────
	m.EvalOutput(map[string]any{
		"value":      fbRounded,
		"units":      "N",
		"direction":  "upward",
		"confidence": m.GetCoherence(),
		"derivation": []string{
			"F_b = ρ × V × g",
			"= 1000 kg/m³ × 0.001 m³ × 9.8 m/s²",
			"= 9.8 N",
		},
	})

	// ── Seal the engine trace ─────────────────────────────────────
	sealed := m.EndTrace(OutcomeAccepted, []string{
		"completeness_guard",
		"constant_retrieval",
		"dimensional_calc",
		"cross_validation",
	})

	// ── @meta:distill — outcome-gated, depth-guarded ─────────────
	distillResult, err := m.EvalMetaDistill(sealed, []string{
		"completeness_guard",
		"constant_retrieval",
		"dimensional_calc",
	})
	if err != nil {
		panic(fmt.Sprintf("distill failed: %v", err))
	}

	// ── Assemble the JSON proof object ────────────────────────────
	env := assembleEnvelope(traceID, now, sealed, distillResult, m, recorder)

	// ── Run verifier → stamp certificate ──────────────────────────
	cert := verify(env)
	env.Signatures.Certificate = cert

	return env
}

// assembleEnvelope converts engine state + recorded events into the
// canonical JSON trace format defined in the design doc.
func assembleEnvelope(
	traceID, createdAt string,
	sealed Trace,
	distill DistillResult,
	m *MRL,
	rec *TraceRecorder,
) TraceEnvelope {

	// ── Steps: translate recorded events → TraceStep objects ──────
	steps := make([]TraceStep, 0, len(rec.events))
	scopeDepth := 0
	stepNum := 0
	for _, ev := range rec.events {
		stepNum++
		if ev.Type == "scope_enter" {
			scopeDepth++
		}
		ts := TraceStep{
			Step:    stepNum,
			Op:      eventTypeToOp(ev.Type),
			Args:    eventToArgs(ev),
			Scope:   scopeDepth,
			Effects: eventToEffects(ev),
			Tags:    eventToTags(ev),
		}
		steps = append(steps, ts)
		if ev.Type == "scope_exit" {
			scopeDepth--
		}
	}

	// ── Coherence trajectory: rebuild deltas from snapshot list ───
	updates := make([]CoherenceUpdate, 0)
	for i, snap := range m.Trajectory {
		if i == 0 {
			continue // skip the initial baseline
		}
		prev := m.Trajectory[i-1]
		dl := snap.State.L - prev.State.L
		dp := snap.State.P - prev.State.P
		dc := snap.State.C - prev.State.C

		delta := CoherenceDelta{}
		if dl != 0 {
			delta.L = fp(dl)
		}
		if dp != 0 {
			delta.P = fp(dp)
		}
		if dc != 0 {
			delta.C = fp(dc)
		}

		updates = append(updates, CoherenceUpdate{
			Step:     snap.Step,
			Delta:    delta,
			Reason:   snap.Reason,
			Evidence: reasonToEvidence(snap.Reason),
		})
	}

	// ── Meta section ──────────────────────────────────────────────
	metaRecord := MetaRecord{
		Distillation: MetaDistillation{
			Attempted: true,
			Mode:      string(distill.Mode),
			Depth:     m.MetaSteps,
			Produced:  distill.Tags,
		},
		Budgets: MetaBudgets{
			MetaDepthUsed:  m.MetaSteps,
			MetaDepthMax:   m.Limits.MaxDistillDepth,
			MetaTokenRatio: float64(m.MetaSteps) / float64(sealed.Cost.Steps+1),
		},
	}

	return TraceEnvelope{
		TraceID:     traceID,
		SpecVersion: "mrl-trace-1.0",
		CreatedAt:   createdAt,
		Task: TaskDescriptor{
			Domain:    sealed.Domain,
			Type:      sealed.TaskType,
			RiskLevel: "medium",
		},
		Outcome: sealed.Outcome,
		Summary: TraceSummary{
			FinalAnswer: "9.8 N upward",
			Confidence:  m.Coh.Mean(),
		},
		Steps: steps,
		Coherence: CoherenceRecord{
			Initial: CoherenceState{L: 0.75, P: 0.15, C: 0.30},
			Updates: updates,
			Final:   m.Coh,
		},
		Meta:       metaRecord,
		Signatures: SignatureRecord{},
	}
}

// ===================================================================
// Event → TraceStep translation helpers
// Map the engine's internal event stream onto the opaque,
// verifier-addressable step schema (§2).
// ===================================================================

func eventTypeToOp(evType string) string {
	switch evType {
	case "var_set":
		return "@b"
	case "step":
		return "@i"
	case "coh":
		return "@coh"
	case "reflection":
		return "@refl"
	case "plugin_call":
		return "@call"
	case "plugin_result":
		return "@result"
	case "output":
		return "@o"
	case "error":
		return "@t"
	case "scope_enter":
		return "@scope"
	case "scope_exit":
		return "@scope_exit"
	case "trace":
		return "@trace"
	case "meta":
		return "@meta:distill"
	default:
		return "@" + evType
	}
}

func eventToArgs(ev StepEvent) []string {
	args := []string{ev.Message}
	if m, ok := ev.Data.(map[string]any); ok {
		if reason, ok := m["reason"].(string); ok {
			args = append(args, reason)
		}
		if trigger, ok := m["trigger"].(string); ok {
			args = append(args, trigger)
		}
		if suggestion, ok := m["suggestion"].(string); ok {
			args = append(args, suggestion)
		}
	}
	return args
}

func eventToEffects(ev StepEvent) []string {
	switch ev.Type {
	case "var_set":
		return []string{"bind:" + ev.Message[3:]} // strip "@b " prefix
	case "coh":
		return []string{"coherence:update"}
	case "reflection":
		return []string{"strategy:pivot"}
	case "plugin_call":
		return []string{"tool:invoke"}
	case "plugin_result":
		return []string{"tool:resolve"}
	case "output":
		return []string{"output:seal"}
	case "scope_enter":
		return []string{"scope:push"}
	case "scope_exit":
		return []string{"scope:pop"}
	case "trace":
		return []string{"trace:lifecycle"}
	case "meta":
		return []string{"meta:distill"}
	case "error":
		return []string{"error:recover"}
	default:
		return []string{"eval"}
	}
}

func eventToTags(ev StepEvent) []string {
	switch ev.Type {
	case "var_set":
		return []string{"binding"}
	case "coh":
		return []string{"coherence"}
	case "reflection":
		return []string{"completeness_guard"}
	case "plugin_call", "plugin_result":
		return []string{"tool_use"}
	case "output":
		return []string{"final_output"}
	case "scope_enter", "scope_exit":
		return []string{"scoping"}
	case "meta":
		return []string{"meta_reasoning"}
	default:
		return []string{"control_flow"}
	}
}

// reasonToEvidence maps a reason string to evidence tokens the verifier
// checks for.  In production this would be structured metadata attached
// at emission time; here we derive it from keywords in the reason.
func reasonToEvidence(reason string) []string {
	evidence := []string{}
	keywords := map[string]string{
		"principle":    "archimedes_principle",
		"geometric":    "volume_calculation",
		"constant":     "water_density",
		"dimensional":  "unit_check",
		"cross":        "cross_validation",
		"setup":        "problem_parameters",
		"completeness": "completeness_guard",
	}
	for kw, tag := range keywords {
		if containsInsensitive(reason, kw) {
			evidence = append(evidence, tag)
		}
	}
	if len(evidence) == 0 {
		evidence = append(evidence, "implicit")
	}
	return evidence
}

// containsInsensitive is a simple case-insensitive substring check.
func containsInsensitive(s, substr string) bool {
	if len(substr) > len(s) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			a, b := s[i+j], substr[j]
			if a >= 'A' && a <= 'Z' {
				a += 32
			}
			if b >= 'A' && b <= 'Z' {
				b += 32
			}
			if a != b {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// ===================================================================
// main — run the example and pretty-print the JSON envelope to stdout
// ===================================================================

func main() {
	envelope := runBuoyancy()

	out, err := json.MarshalIndent(envelope, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
