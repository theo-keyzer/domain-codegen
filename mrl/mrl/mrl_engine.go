package main

import (
	"fmt"
	"math"
	"time"
)

// ===================================================================
// Core Types
// ===================================================================

// Value represents any MRL value type.
type Value any

// Env is a single scope frame: variable name → value.
// Variables are immutable once set within a frame (v1.1 §3.1).
type Env map[string]Value

// StepEvent represents a streaming event during MRL execution.
type StepEvent struct {
	SessionID string `json:"session_id"`
	Seq       int    `json:"seq"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
}

// Streamer interface for emitting execution events.
type Streamer interface {
	Emit(event StepEvent)
}

// Plugin is a function that can be called from MRL.
type Plugin func(args []Value) (Value, error)

// ===================================================================
// v1.1 §5 — Coherence: 3-Dimension State & Trajectory
// ===================================================================

// CoherenceState holds the three canonical dimensions defined in §5.2.
//
//	L — Logical consistency   (0–1)
//	P — Progress toward goal  (0–1)
//	C — Completeness          (0–1)
type CoherenceState struct {
	L float64 `json:"l"`
	P float64 `json:"p"`
	C float64 `json:"c"`
}

// Mean returns the arithmetic mean of all three dimensions.
// This is the single scalar the engine exposes for backward-compatible
// reads (e.g. GetCoherence).
func (cs CoherenceState) Mean() float64 {
	return (cs.L + cs.P + cs.C) / 3.0
}

// CoherenceSnapshot captures a full state at a numbered step.
// The slice of these forms the trajectory required by §6.1.
type CoherenceSnapshot struct {
	Step   int           `json:"step"`
	State  CoherenceState `json:"state"`
	Reason string        `json:"reason,omitempty"`
}

// CoherenceDeltas carries the signed adjustments passed to @coh.
// Zero-value fields mean "no change to this dimension".
type CoherenceDeltas struct {
	L      float64 `json:"l,omitempty"`
	P      float64 `json:"p,omitempty"`
	C      float64 `json:"c,omitempty"`
	Reason string  `json:"reason"`
}

// clamp restricts v to [0, 1].
func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// ===================================================================
// v1.1 §3.1 — Scope Stack (Patch 1: Lexical Scope & Immutability)
// ===================================================================

// ScopeStack is an ordered list of Env frames.  The last element is the
// innermost (current) scope.  Lookup walks from inner to outer.
// SetVar only writes to the current (innermost) frame and panics if the
// name already exists in that frame — enforcing single-assignment per scope.
type ScopeStack struct {
	frames []Env
}

// NewScopeStack creates a stack with one root frame.
func NewScopeStack() *ScopeStack {
	return &ScopeStack{
		frames: []Env{make(Env)},
	}
}

// Push opens a new lexical scope (@scope enter).
func (ss *ScopeStack) Push() {
	ss.frames = append(ss.frames, make(Env))
}

// Pop closes the innermost lexical scope (@scope exit).
// Panics if the caller tries to pop the root frame.
func (ss *ScopeStack) Pop() {
	if len(ss.frames) <= 1 {
		panic("cannot pop root scope")
	}
	ss.frames = ss.frames[:len(ss.frames)-1]
}

// current returns the innermost frame.
func (ss *ScopeStack) current() Env {
	return ss.frames[len(ss.frames)-1]
}

// Set binds name → value in the current frame.
// If name already exists in THIS frame, it panics with REBIND_NOT_ALLOWED
// (§3.1 normative rule).  A name that exists only in an outer frame is
// fine — that constitutes legal shadowing inside an explicit @scope.
func (ss *ScopeStack) Set(name string, value Value) {
	frame := ss.current()
	if _, exists := frame[name]; exists {
		panic(fmt.Sprintf("REBIND_NOT_ALLOWED: variable '%s' is immutable in current scope", name))
	}
	frame[name] = value
}

// Get resolves name by walking frames from innermost to outermost.
// Returns the value and true if found; zero-Value and false otherwise.
func (ss *ScopeStack) Get(name string) (Value, bool) {
	for i := len(ss.frames) - 1; i >= 0; i-- {
		if val, ok := ss.frames[i][name]; ok {
			return val, true
		}
	}
	return nil, false
}

// Depth returns how many frames are on the stack (1 = root only).
func (ss *ScopeStack) Depth() int {
	return len(ss.frames)
}

// ===================================================================
// v1.1 §3.2 — Meta Limits (Patch 2: Circuit Breaker)
// ===================================================================

// MetaLimits encodes the global invariants from §3.2.
type MetaLimits struct {
	MaxDistillDepth   int     `json:"max_distill_depth"`   // hard ceiling on nested meta-ops
	MaxMetaTokenRatio float64 `json:"max_meta_tokens_ratio"` // fraction of context budget
}

// DefaultMetaLimits returns the values mandated by the spec.
func DefaultMetaLimits() MetaLimits {
	return MetaLimits{
		MaxDistillDepth:   3,
		MaxMetaTokenRatio: 0.25,
	}
}

// ===================================================================
// v1.1 §3.3 — Trace & Distillation Types (Patch 3: Outcome Gate)
// ===================================================================

// TraceOutcome is the finite set of allowed outcome labels.
// Only "accepted" permits promotion to reusable_strategy (§3.3).
type TraceOutcome string

const (
	OutcomeAccepted TraceOutcome = "accepted"
	OutcomeRejected TraceOutcome = "rejected"
	OutcomeFailed   TraceOutcome = "failed"
)

// DistillMode is the bimodal output routing mandated by §3.3.
type DistillMode string

const (
	DistillSuccess DistillMode = "success" // outcome == accepted  → reusable_strategy
	DistillFailure DistillMode = "failure" // outcome != accepted  → anti_pattern
)

// DistillResult is the output of a @meta:distill operation.
type DistillResult struct {
	Mode    DistillMode  `json:"mode"`
	Pattern string       `json:"pattern"`     // e.g. "reusable_strategy" | "anti_pattern"
	TraceID string       `json:"trace_id"`
	Tags    []string     `json:"tags,omitempty"`
}

// Trace is the full structured record of one reasoning episode (§6.1).
type Trace struct {
	ID                string              `json:"id"`
	Domain            string              `json:"domain"`
	TaskType          string              `json:"task_type"`
	Outcome           TraceOutcome        `json:"outcome"`
	CoherenceTrajectory []CoherenceSnapshot `json:"coherence_trajectory"`
	StrategyTags      []string            `json:"strategy_tags,omitempty"`
	Cost              TraceCost           `json:"cost"`
	StartedAt         time.Time           `json:"started_at"`
	FinishedAt        time.Time           `json:"finished_at,omitempty"`
}

// TraceCost tracks resource consumption for efficiency benchmarking (§6.1).
type TraceCost struct {
	Steps     int `json:"steps"`
	PluginCalls int `json:"plugin_calls"`
}

// ===================================================================
// MRL Engine — v1.1
// ===================================================================

// MRL is the core execution engine, upgraded to v1.1.
type MRL struct {
	// Scoped variable store — replaces the flat Env map (§3.1).
	Scopes *ScopeStack

	Plugins   map[string]Plugin
	Streamer  Streamer
	SessionID string
	Seq       int
	Output    Value

	// --- Coherence (§5) -------------------------------------------
	Coh        CoherenceState       // live 3-dimension state
	Trajectory []CoherenceSnapshot  // append-only history

	// --- Meta circuit breaker (§3.2) --------------------------------
	MetaDepth  int        // current nesting depth of meta-operations
	MetaSteps  int        // total meta-op steps consumed (for ratio)
	Limits     MetaLimits // hard limits

	// --- Active trace recording (§6.1) -----------------------------
	ActiveTrace *Trace
}

// NewMRL creates a new MRL engine instance with v1.1 defaults.
func NewMRL(streamer Streamer) *MRL {
	initial := CoherenceState{L: 0.5, P: 0.0, C: 0.0}
	m := &MRL{
		Scopes:    NewScopeStack(),
		Plugins:   make(map[string]Plugin),
		Streamer:  streamer,
		SessionID: fmt.Sprintf("session_%d", time.Now().UnixMilli()),
		Coh:       initial,
		Trajectory: []CoherenceSnapshot{
			{Step: 0, State: initial, Reason: "initial"},
		},
		Limits: DefaultMetaLimits(),
	}
	return m
}

// step emits a streaming event.
func (m *MRL) step(eventType, msg string, data any) {
	m.Seq++
	if m.ActiveTrace != nil {
		m.ActiveTrace.Cost.Steps++
	}
	m.Streamer.Emit(StepEvent{
		SessionID: m.SessionID,
		Type:      eventType,
		Seq:       m.Seq,
		Message:   msg,
		Data:      data,
	})
}

// ===================================================================
// v1.1 §3.1 — @b (Bind) & @scope
// ===================================================================

// SetVar binds a variable in the current scope.
// Panics with REBIND_NOT_ALLOWED if the name already exists in this scope.
func (m *MRL) SetVar(name string, value Value) {
	m.Scopes.Set(name, value) // immutability enforced inside ScopeStack.Set
	m.step("var_set", fmt.Sprintf("@b %s", name), value)
}

// GetVar resolves a variable by walking scopes inner→outer.
func (m *MRL) GetVar(name string) Value {
	val, ok := m.Scopes.Get(name)
	if !ok {
		panic(fmt.Sprintf("variable not found: %s", name))
	}
	return val
}

// EvalScope implements @scope — pushes a new frame, runs the body,
// then pops the frame.  Any bindings made inside are invisible after return.
// This is the ONLY legal way to shadow a variable (§3.1).
func (m *MRL) EvalScope(body func()) {
	m.step("scope_enter", "@scope enter", nil)
	m.Scopes.Push()
	defer func() {
		m.Scopes.Pop()
		m.step("scope_exit", "@scope exit", nil)
	}()
	body()
}

// ===================================================================
// v1.1 §5 — @coh (Coherence Update)
// ===================================================================

// UpdateCoherence applies a delta set to the live coherence state,
// clamps every dimension to [0,1], and appends a trajectory snapshot.
// This is the engine-level implementation of @coh.
func (m *MRL) UpdateCoherence(d CoherenceDeltas) {
	m.Coh.L = clamp(m.Coh.L + d.L)
	m.Coh.P = clamp(m.Coh.P + d.P)
	m.Coh.C = clamp(m.Coh.C + d.C)

	snap := CoherenceSnapshot{
		Step:   m.Seq,
		State:  m.Coh,
		Reason: d.Reason,
	}
	m.Trajectory = append(m.Trajectory, snap)

	m.step("coh", "@coh update", map[string]any{
		"state":  m.Coh,
		"delta":  d,
		"reason": d.Reason,
	})
}

// EvalCoh is the high-level @coh primitive: accepts deltas and reason, delegates.
func (m *MRL) EvalCoh(l, p, c float64, reason string) {
	m.UpdateCoherence(CoherenceDeltas{L: l, P: p, C: c, Reason: reason})
}

// GetCoherence returns the scalar mean for backward compatibility.
func (m *MRL) GetCoherence() float64 {
	return m.Coh.Mean()
}

// ===================================================================
// v1.1 §5 — @refl (Reflection)
// ===================================================================

// EvalRefl implements @refl — emits a reflection event with a trigger
// label and a suggestion.  In a full system the suggestion would be
// routed back to the LLM; here it becomes a first-class event on the
// stream so downstream consumers (verifier, UI) can act on it.
func (m *MRL) EvalRefl(trigger, suggestion string) {
	m.step("reflection", "@refl", map[string]any{
		"trigger":    trigger,
		"suggestion": suggestion,
		"coh":        m.Coh,
	})
}

// ===================================================================
// v1.1 §2 — @i (If), @t (Try), @g (Get), @o (Output)  [unchanged API, updated internals]
// ===================================================================

// EvalIf implements @i — conditional execution.
// Coherence is no longer bumped unconditionally on "true"; callers
// should use EvalCoh explicitly if the branch warrants a coherence change.
func (m *MRL) EvalIf(cond bool, thenFn, elseFn func()) {
	m.step("step", "@i evaluating condition", cond)
	if cond {
		thenFn()
	} else if elseFn != nil {
		elseFn()
	}
}

// EvalTry implements @t — try with fallback on panic.
func (m *MRL) EvalTry(mainFn, altFn func()) {
	defer func() {
		if r := recover(); r != nil {
			m.step("error", "@t recovered from failure", r)
			// Record the failure as a negative coherence event.
			m.UpdateCoherence(CoherenceDeltas{
				L:      -0.1,
				Reason: fmt.Sprintf("try-block panic: %v", r),
			})
			if altFn != nil {
				altFn()
			}
		}
	}()
	mainFn()
}

// EvalGet implements @g — structure access with default support.
func (m *MRL) EvalGet(obj any, key string) Value {
	m.step("step", "@g accessing structure", key)

	switch v := obj.(type) {
	case map[string]any:
		if val, ok := v[key]; ok {
			return val
		}
		panic(fmt.Sprintf("key not found: %s", key))
	case map[string]Value:
		if val, ok := v[key]; ok {
			return val
		}
		panic(fmt.Sprintf("key not found: %s", key))
	default:
		panic("unsupported structure type for @g")
	}
}

// EvalGetWithDefault is @g with a fallback value — avoids panics for optional fields.
func (m *MRL) EvalGetWithDefault(obj any, key string, defaultVal Value) Value {
	m.step("step", "@g accessing structure (with default)", key)

	switch v := obj.(type) {
	case map[string]any:
		if val, ok := v[key]; ok {
			return val
		}
		return defaultVal
	case map[string]Value:
		if val, ok := v[key]; ok {
			return val
		}
		return defaultVal
	default:
		return defaultVal
	}
}

// EvalOutput implements @o — seals the output and emits the final event.
func (m *MRL) EvalOutput(v Value) {
	m.Output = v
	m.step("output", "@o final answer produced", v)
}

// ===================================================================
// Plugin Registry
// ===================================================================

// RegisterPlugin adds a plugin to the engine.
func (m *MRL) RegisterPlugin(name string, fn Plugin) {
	m.Plugins[name] = fn
}

// CallPlugin executes a registered plugin and records the call in the trace.
func (m *MRL) CallPlugin(name string, args ...Value) Value {
	m.step("plugin_call", name, args)

	plugin, ok := m.Plugins[name]
	if !ok {
		panic(fmt.Sprintf("plugin not found: %s", name))
	}

	result, err := plugin(args)
	if err != nil {
		panic(fmt.Sprintf("plugin error in %s: %v", name, err))
	}

	if m.ActiveTrace != nil {
		m.ActiveTrace.Cost.PluginCalls++
	}

	m.step("plugin_result", name, result)
	return result
}

// ===================================================================
// v1.1 §3.2 — @meta:distill with Circuit Breaker
// ===================================================================

// EvalMetaDistill implements @meta:distill with the depth guard and
// outcome gate mandated by §3.2 and §3.3.
//
// It returns a DistillResult whose Mode is determined solely by the
// trace's Outcome field — never by coherence alone.
func (m *MRL) EvalMetaDistill(trace Trace, tags []string) (DistillResult, error) {
	// --- Circuit breaker (§3.2) -------------------------------------
	m.MetaDepth++
	defer func() { m.MetaDepth-- }()

	if m.MetaDepth > m.Limits.MaxDistillDepth {
		return DistillResult{}, fmt.Errorf("META_DEPTH_LIMIT_EXCEEDED: depth %d > max %d",
			m.MetaDepth, m.Limits.MaxDistillDepth)
	}

	m.MetaSteps++

	m.step("meta", "@meta:distill enter", map[string]any{
		"trace_id":   trace.ID,
		"depth":      m.MetaDepth,
		"outcome":    trace.Outcome,
	})

	// --- Outcome gate (§3.3) ----------------------------------------
	var result DistillResult

	switch {
	case trace.Outcome == OutcomeAccepted:
		// Success path → reusable strategy, eligible for promotion.
		result = DistillResult{
			Mode:    DistillSuccess,
			Pattern: "reusable_strategy",
			TraceID: trace.ID,
			Tags:    tags,
		}
	default:
		// Failure path → anti-pattern only.  Never promoted.
		result = DistillResult{
			Mode:    DistillFailure,
			Pattern: "anti_pattern",
			TraceID: trace.ID,
			Tags:    tags,
		}
	}

	m.step("meta", "@meta:distill complete", result)
	return result, nil
}

// ===================================================================
// v1.1 §6.1 — Trace Lifecycle
// ===================================================================

// BeginTrace opens a new trace recording session.
func (m *MRL) BeginTrace(id, domain, taskType string) {
	m.ActiveTrace = &Trace{
		ID:        id,
		Domain:    domain,
		TaskType:  taskType,
		StartedAt: time.Now(),
	}
	m.step("trace", "@trace begin", map[string]any{
		"id":        id,
		"domain":    domain,
		"task_type": taskType,
	})
}

// EndTrace seals the active trace with an outcome and snapshots the
// final coherence trajectory.  Returns the completed Trace.
func (m *MRL) EndTrace(outcome TraceOutcome, strategyTags []string) Trace {
	if m.ActiveTrace == nil {
		panic("EndTrace called with no active trace")
	}

	m.ActiveTrace.Outcome = outcome
	m.ActiveTrace.StrategyTags = strategyTags
	m.ActiveTrace.CoherenceTrajectory = m.Trajectory
	m.ActiveTrace.FinishedAt = time.Now()

	sealed := *m.ActiveTrace
	m.ActiveTrace = nil

	m.step("trace", "@trace sealed", map[string]any{
		"id":      sealed.ID,
		"outcome": sealed.Outcome,
		"cost":    sealed.Cost,
	})

	return sealed
}

// ===================================================================
// Built-in Plugins
// ===================================================================

// MathAdd adds two numbers.
func MathAdd(args []Value) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("math:add requires exactly 2 arguments")
	}
	a, ok1 := toFloat64(args[0])
	b, ok2 := toFloat64(args[1])
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("math:add requires numeric arguments")
	}
	return a + b, nil
}

// MathSubtract subtracts the second argument from the first.
func MathSubtract(args []Value) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("math:subtract requires exactly 2 arguments")
	}
	a, ok1 := toFloat64(args[0])
	b, ok2 := toFloat64(args[1])
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("math:subtract requires numeric arguments")
	}
	return a - b, nil
}

// MathMultiply multiplies all arguments together (variadic).
// The buoyancy trace calls @math:multiply with 3 args (ρ × V × g);
// this implementation handles any N ≥ 1.
func MathMultiply(args []Value) (Value, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("math:multiply requires at least 1 argument")
	}
	product := 1.0
	for i, arg := range args {
		v, ok := toFloat64(arg)
		if !ok {
			return nil, fmt.Errorf("math:multiply: argument %d is not numeric", i)
		}
		product *= v
	}
	return product, nil
}

// MathPow raises base to the power of exp.
// Used by the buoyancy trace: @math:pow 0.1 3 → 0.001
func MathPow(args []Value) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("math:pow requires exactly 2 arguments (base, exp)")
	}
	base, ok1 := toFloat64(args[0])
	exp, ok2 := toFloat64(args[1])
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("math:pow requires numeric arguments")
	}
	return math.Pow(base, exp), nil
}

// MathRound rounds a number to the specified number of decimal places.
func MathRound(args []Value) (Value, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("math:round requires exactly 2 arguments (value, decimals)")
	}
	val, ok1 := toFloat64(args[0])
	dec, ok2 := toFloat64(args[1])
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("math:round requires numeric arguments")
	}
	shift := math.Pow(10, dec)
	return math.Round(val*shift) / shift, nil
}

// TimeNow returns the current timestamp as an RFC 3339 string.
func TimeNow(args []Value) (Value, error) {
	return time.Now().Format(time.RFC3339), nil
}

// TimeUnix returns the current Unix timestamp.
func TimeUnix(args []Value) (Value, error) {
	return time.Now().Unix(), nil
}

// WeatherMock returns mock weather data (demo / test plugin).
func WeatherMock(args []Value) (Value, error) {
	location := "Unknown"
	if len(args) > 0 {
		location = fmt.Sprintf("%v", args[0])
	}
	return map[string]any{
		"location": location,
		"temp":     12,
		"unit":     "C",
		"cond":     "Cloudy",
		"humidity": 65,
	}, nil
}

// StringConcat concatenates all arguments as strings.
func StringConcat(args []Value) (Value, error) {
	result := ""
	for _, arg := range args {
		result += fmt.Sprintf("%v", arg)
	}
	return result, nil
}

// LogMessage prints a debug message to stdout.
func LogMessage(args []Value) (Value, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("log requires at least 1 argument")
	}
	message := fmt.Sprintf("%v", args[0])
	fmt.Println("[MRL LOG]", message)
	return message, nil
}

// ===================================================================
// Helper Functions
// ===================================================================

// toFloat64 converts various numeric types to float64.
func toFloat64(v Value) (float64, bool) {
	switch val := v.(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int:
		return float64(val), true
	case int32:
		return float64(val), true
	case int64:
		return float64(val), true
	default:
		return 0, false
	}
}

// RegisterDefaultPlugins registers all built-in plugins.
func (m *MRL) RegisterDefaultPlugins() {
	m.RegisterPlugin("math:add", MathAdd)
	m.RegisterPlugin("math:subtract", MathSubtract)
	m.RegisterPlugin("math:multiply", MathMultiply)
	m.RegisterPlugin("math:pow", MathPow)
	m.RegisterPlugin("math:round", MathRound)
	m.RegisterPlugin("time:now", TimeNow)
	m.RegisterPlugin("time:unix", TimeUnix)
	m.RegisterPlugin("weather", WeatherMock)
	m.RegisterPlugin("string:concat", StringConcat)
	m.RegisterPlugin("log", LogMessage)
}
