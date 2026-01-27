# MRL Meta-Reasoner: Complete Implementation Guide
**Self-Improving Reasoning Systems Through Trace Libraries**  
Version 2026.1 | Last Updated: January 27, 2026

---

## Table of Contents
1. [What is a Meta-Reasoner?](#what-is-a-meta-reasoner)
2. [Core Architecture](#core-architecture)
3. [Operational Modes](#operational-modes)
4. [Complete Example: Physics Reasoning](#complete-example-physics-reasoning)
5. [Failure Taxonomies](#failure-taxonomies)
6. [Implementation Patterns](#implementation-patterns)
7. [Cross-Domain Transfer](#cross-domain-transfer)
8. [Heuristic Lifecycle Management](#heuristic-lifecycle-management)
9. [Quick Reference](#quick-reference)

---

## What is a Meta-Reasoner?

### Definition

A **meta-reasoner** is a higher-level supervisory system that operates over a library of MRL reasoning traces. It treats past problem-solving experiences as **reusable strategy memory**, enabling:

- **Meta-cognitive reuse**: Learning from past successful/failed reasoning patterns
- **Reflection distillation**: Extracting generalizable strategies from specific traces
- **Strategy composition**: Combining proven patterns for novel problems
- **Self-improvement**: Iteratively refining reasoning approaches

### What It's NOT

❌ **Not a planner** - Goes beyond step sequencing  
❌ **Not RAG** - Retrieves reasoning strategies, not just facts  
❌ **Not prompt chaining** - Learns from execution, not templates

### What It IS

✅ **Strategy-induction system** - Learns how to reason  
✅ **Operating on reasoning trajectories** - Uses complete execution traces  
✅ **Coherence-guided** - Uses coherence as universal reward signal

**Related to:**
- AlphaZero-style self-play (but symbolic)
- Program synthesis from execution traces
- Proof tactic learning in theorem provers

---

## Core Architecture

### 1. Trace Library Structure

Each trace is stored as a rich, structured object with metadata:

```json
{
  "id": "trace-20260127-buoyant-001",
  "domain": "classical physics / fluids",
  "task_type": "compute derived quantity + constants lookup",
  "outcome": "correct",
  
  "coherence": {
    "initial": {"l": 0.70, "p": 0.10, "c": 0.30},
    "final": {"l": 0.95, "p": 0.90, "c": 0.95},
    "trajectory": [
      {"step": 1, "l": 0.70, "p": 0.10, "c": 0.30},
      {"step": 2, "l": 0.80, "p": 0.30, "c": 0.45},
      {"step": 3, "l": 0.83, "p": 0.30, "c": 0.70}
    ],
    "gain": 0.40
  },
  
  "strategy_tags": [
    "early completeness check",
    "missing-constant reflection",
    "@coh guardrail",
    "Archimedes primitive"
  ],
  
  "patterns_used": {
    "reflection_triggers": ["@i (< ?state.coh.c 0.65) → @refl"],
    "coherence_updates": ["@coh ?state c:+0.22 l:+0.03 after retrieving rho"],
    "tool_sequences": ["@math:pow → @math:multiply → @math:round"]
  },
  
  "cost": {
    "tokens": 892,
    "steps": 7,
    "tool_calls": 3,
    "time_ms": 1240
  },
  
  "key_insights": [
    "Completeness check before computation prevents errors",
    "Unit consistency verification catches dimensional mistakes"
  ],
  
  "mrl_trace": "..." // Complete MRL execution
}
```

### 2. Metadata Enrichment

Traces are automatically enriched with:

| Field | Purpose | Extraction Method |
|-------|---------|-------------------|
| **domain** | Task categorization | NLP classification or manual tagging |
| **outcome** | Success/failure/partial | Final coherence + goal achievement |
| **coherence_trajectory** | Learning signal | Track `@coh` updates through execution |
| **strategy_tags** | Pattern indexing | Extract from primitives + reflections |
| **failure_mode** | Error classification | Analyze coherence collapse points |
| **cost** | Efficiency metrics | Token counting + timing |

### 3. Three Core Operations

```lisp
; Retrieve similar traces
(@meta:retrieve 
  query:"physics buoyancy"
  domain:"classical mechanics"
  min_coh:0.85
  → ?candidates)

; Compose strategy from multiple traces
(@meta:compose
  patterns:["completeness_guard" "unit_check"]
  problem:?new_task
  → ?strategy_prefix)

; Distill new heuristics from failures
(@meta:distill
  failure:?trace
  focus:"coherence_collapse"
  → ?new_heuristic)
```

---

## Operational Modes

### Mode 1: Retrieval + Reuse

**When:** New problem resembles past traces  
**Goal:** Inject proven strategy skeleton

```lisp
;; Meta-phase (before normal reasoning)
@recall domain:"physics" tags:["buoyancy" "density"] min_coh:0.85 → ?candidates
@rank ?candidates by:"coh.l * progress + recency" → ?best_trace

;; Extract and inject strategy prefix
@b ?strategy_prefix (extract ?best_trace 
  "reflection on missing constants"
  "Archimedes application")
@prepend ?strategy_prefix

;; Normal reasoning now has bias toward proven patterns
@b ?goal "Net force on floating 0.2 m³ wood block density 650 kg/m³"
@b ?state {coh: {l:0.70 p:0.10 c:0.30}}

;; Injected prefix → early completeness check fires automatically
@i (< ?state.coh.c 0.60)
   (@refl suggestion:"Retrieve wood & water densities + check submersion")
```

**Benefits:**
- Bootstraps new traces with proven guardrails
- Reduces token overhead (no re-discovery)
- Maintains coherence from step 1

### Mode 2: Pattern Composition

**When:** No perfect match, but pieces fit  
**Goal:** Stitch together sub-strategies

```lisp
;; Problem: Debug memory leak in distributed cache
@recall tags:["debugging" "memory"] → ?debug_traces
@recall tags:["distributed_systems" "cache"] → ?arch_traces

;; Extract complementary patterns
@extract ?debug_traces pattern:"snapshot_diff" → ?p1
@extract ?arch_traces pattern:"node_comparison" → ?p2

;; Compose into unified strategy
@compose from:[?p1 ?p2] goal:"isolate leak source" → ?strategy

;; Injected composite strategy
@b ?strategy {
  step1: "Take heap snapshots on all nodes"
  step2: "Diff snapshots across time and space"
  step3: "Correlate with cache operation logs"
  reflection_points: [
    "@i (diff_same_across_nodes) → single-node issue",
    "@i (diff_different_per_node) → distributed issue"
  ]
}
```

**Benefits:**
- Handles novel combinations
- Leverages partial matches
- Creates new patterns through fusion

### Mode 3: Failure-Directed Learning

**When:** Coherence drops or loops detected  
**Goal:** Extract negative lessons

```lisp
;; Detect failure pattern
@analyze ?failing_trace → {
  coherence_collapse: {
    step: 4,
    trigger: "attempted enumeration without feedback",
    l_drop: -0.45,
    p_drop: -0.30
  },
  root_cause: "epistemic_mismatch",
  category: "category_error"
}

;; Distill corrective heuristic
@distill ?analysis → ?heuristic {
  pattern: "hidden_state_without_observation",
  condition: "no operation yields new information",
  action: "@refl suggestion:'Prove impossibility instead of constructing solution'",
  anti_pattern: "state_enumeration",
  domains: ["logic_puzzles" "distributed_systems" "incomplete_info"]
}

;; Store for future use
@store ?heuristic 
  tags:["epistemic_guard" "impossibility_proof"]
  confidence:0.78
  version:"v1.0"
```

**Benefits:**
- Learns what NOT to do
- Generalizes across domains
- Prevents repeated failures

---

## Complete Example: Physics Reasoning

### Problem: Buoyant Force on Submerged Cube

```lisp
(trace "buoyant force on submerged cube")

;; --- Initial State ---
@b ?goal "Calculate buoyant force on fully submerged 10cm cube in water"

@b ?state {
  beliefs: [
    "Cube side length = 0.1 m",
    "Cube fully submerged in water",
    "Acceleration due to gravity g ≈ 9.8 m/s²"
  ]
  assumptions: ["Water is pure at 4°C", "Cube is rigid and non-porous"]
  contradictions: []
  coherence: {
    logical: 0.70      ; Initial setup makes sense
    progress: 0.10     ; Just starting
    completeness: 0.30 ; Missing key elements
    consistency: 0.90  ; No contradictions yet
  }
  uncertainty: ["Exact water density?", "Any surface tension effects?"]
}

;; --- Step 1: Retrieve Fundamental Principle ---
@b ?archimedes {
  statement: "Buoyant force equals weight of displaced fluid"
  formula: "F_b = ρ_fluid × V_displaced × g"
  source: "Archimedes' principle (c. 246 BC)"
  confidence: 0.99
}

@b ?state.beliefs += ?archimedes.statement

(@coh ?state 
  logical: +0.10   ; Added fundamental principle
  progress: +0.20  ; Clear path forward
  completeness: +0.15
  reason: "Retrieved governing physical law")

;; --- Step 2: Calculate Displaced Volume ---
@t (@math:pow 0.1 3 → ?volume) 
   alt:(@refl ?state suggestion:"Check cube dimension units")
   
@b ?volume_str "V = (0.1 m)³ = 0.001 m³ = 1 liter"
@b ?state.beliefs += ?volume_str

(@coh ?state
  logical: +0.05
  progress: +0.15
  consistency: +0.05  ; Calculation matches dimensional analysis
  reason: "Volume calculation consistent with geometry")

;; --- Step 3: Coherence-Triggered Reflection ---
@i (< ?state.coherence.completeness 0.60)
   (@refl ?state 
          trigger:"incomplete_parameters"
          suggestion:"Need fluid density for buoyancy calculation")
   
   ;; Show multiple options for density
   @b ?density_options {
     "Pure water at 4°C": 1000       ; kg/m³
     "Sea water": 1025
     "Lake water": 998
   }
   
   @s focus:"Assuming pure water unless specified otherwise"
   @b ?rho_selected 1000

@b ?state.beliefs += "Using water density ρ = 1000 kg/m³"
@b ?state.assumptions += "Pure water at standard conditions"

(@coh ?state
  completeness: +0.25
  logical: +0.03
  reason: "Added required parameter with noted assumption")

;; --- Step 4: Calculate Buoyant Force ---
@b ?Fb_raw (@math:multiply ?rho_selected ?volume 9.8)
@b ?Fb_rounded (@math:round ?Fb_raw 2)  ; 9.8 N

@b ?derivation {
  step1: "F_b = ρ × V × g"
  step2: "= 1000 kg/m³ × 0.001 m³ × 9.8 m/s²"
  step3: "= 1 kg × 9.8 m/s²"
  step4: "= 9.8 N"
}

@b ?state.beliefs += ?derivation

(@coh ?state
  logical: +0.08
  progress: +0.30
  consistency: +0.10  ; Units check: (kg/m³)×m³×(m/s²) = N
  reason: "Complete derivation with unit consistency")

;; --- Step 5: Cross-Validation ---
@i (debug_mode:true)
   ;; Alternative method: Weight of displaced water
   @b ?mass_water (@math:multiply ?rho_selected ?volume)  ; 1 kg
   @b ?weight_water (@math:multiply ?mass_water 9.8)      ; 9.8 N
   
   @if (≈ ?Fb_rounded ?weight_water 0.01)
        (@coh ?state 
              consistency: +0.15
              reason: "Two independent calculations agree")
   @else
        (@coh ?state logical: -0.30 reason: "Calculation discrepancy")
        (@refl ?state trigger:"calculation_mismatch"
               suggestion:"Re-examine volume or density")

;; --- Step 6: Edge Case Analysis ---
@b ?edge_cases {
  "Partial submersion": "Formula → F_b = ρ g V_submerged"
  "Different fluid": "Use appropriate density"
  "Accelerating container": "Effective g changes"
  "Surface tension": "Negligible for macroscopic cube"
}

@b ?state.beliefs += "Assumptions: full submersion, static fluid"

(@coh ?state
  completeness: +0.10
  reason: "Acknowledged limitations and boundary conditions")

;; --- Step 7: Output with Confidence ---
@o {
  problem: "Buoyant force on 10cm cube in water"
  solution: {
    value: "9.8 N"
    units: "newtons"
    direction: "Upward"
  }
  derivation: ?derivation
  assumptions: ?state.assumptions
  confidence: {
    overall: (@math:mean ?state.coherence)  ; ≈ 0.88
    breakdown: ?state.coherence
    factors: [
      "High confidence in Archimedes' principle",
      "Standard density value used",
      "Unit consistency verified",
      "Cross-validated with alternative method"
    ]
  }
  related_concepts: [
    "Specific gravity = ρ_object/ρ_fluid",
    "If F_b > weight, object floats",
    "For this cube, weight must exceed 9.8 N to sink"
  ]
}
```

### What Gets Stored in the Library

```json
{
  "id": "trace-buoyant-001",
  "domain": "classical_physics_fluids",
  "task_type": "derived_quantity_calculation",
  "outcome": "correct",
  
  "extracted_patterns": [
    {
      "name": "completeness_guard",
      "trigger": "@i (< ?state.coh.c 0.60)",
      "action": "@refl suggestion:'identify missing parameters'",
      "effectiveness": 0.95
    },
    {
      "name": "unit_consistency_check",
      "trigger": "after_calculation",
      "action": "dimensional_analysis",
      "effectiveness": 0.92
    },
    {
      "name": "cross_validation",
      "trigger": "critical_result",
      "action": "alternative_method_comparison",
      "effectiveness": 0.89
    }
  ],
  
  "reusable_for": [
    "hydrostatics",
    "density_calculations",
    "force_derivations",
    "constant_lookup_tasks"
  ]
}
```

---

## Failure Taxonomies

### Critical Categories

```lisp
;; Failure taxonomy schema
@b ?failure_types {
  
  ;; 1. Category Error
  category_error: {
    description: "Wrong approach for problem type"
    examples: [
      "Attempting constructive solution to impossibility proof",
      "Using numeric methods on symbolic problem"
    ]
    detection: "l drops while p increases"
    correction: "Goal rewrite or domain shift"
  }
  
  ;; 2. Missing Observability
  epistemic_mismatch: {
    description: "Insufficient information to proceed"
    examples: [
      "Hidden state without feedback channel",
      "Underconstrained system"
    ]
    detection: "c stalls below 0.6, p oscillates"
    correction: "Add information-gathering step or prove impossibility"
  }
  
  ;; 3. False Completeness
  premature_closure: {
    description: "Stopped before full solution"
    examples: [
      "Assumed constant without verification",
      "Ignored edge cases"
    ]
    detection: "High c but low consistency"
    correction: "Enforce edge case enumeration"
  }
  
  ;; 4. Tool Misuse
  tool_error: {
    description: "Wrong tool for task"
    examples: [
      "Web search for mathematical proof",
      "Code execution for conceptual question"
    ]
    detection: "Repeated tool failures, no coherence gain"
    correction: "Strategy replacement"
  }
  
  ;; 5. Logical Inconsistency
  coherence_collapse: {
    description: "Contradictory beliefs accumulated"
    examples: [
      "Circular reasoning",
      "Conflicting assumptions"
    ]
    detection: "l < 0.3, consistency near 0"
    correction: "Backtrack to last coherent state"
  }
}
```

### Failure Detection Patterns

```lisp
;; Automated failure detection
@foreach ?trace ?library
  
  ;; Category error detection
  @i (and (< ?trace.final.l 0.4) 
          (> ?trace.max_progress 0.7)
          (dropped_then_recovered ?trace.coherence.p))
     (@tag ?trace failure_mode:"category_error")
  
  ;; Epistemic mismatch
  @i (and (stalled ?trace.coherence.c below:0.6)
          (oscillates ?trace.coherence.p))
     (@tag ?trace failure_mode:"epistemic_mismatch")
  
  ;; Premature closure
  @i (and (> ?trace.final.c 0.85)
          (< ?trace.final.consistency 0.6))
     (@tag ?trace failure_mode:"premature_closure")
  
  ;; Tool misuse
  @i (repeated_tool_failures ?trace count:>3)
     (@tag ?trace failure_mode:"tool_error")
```

---

## Implementation Patterns

### Pattern 1: Early Reflection Guards

**Problem:** Proceeding with incomplete information  
**Solution:** Coherence-based early stopping

```lisp
;; Always check completeness before critical computations
@i (< ?state.coh.c 0.65)
   (@refl ?state 
          trigger:"pre_computation_check"
          suggestion:"Enumerate missing parameters")
   @halt_until_complete
```

**When to use:**
- Physics derivations (need all constants)
- Database queries (need all constraints)
- API calls (need all required params)

### Pattern 2: Cross-Validation

**Problem:** Single-path errors go undetected  
**Solution:** Multiple independent derivations

```lisp
;; Primary calculation
@b ?result_primary (@method_A ?inputs)

;; Alternative calculation
@b ?result_alternative (@method_B ?inputs)

;; Consistency check
@i (≈ ?result_primary ?result_alternative tolerance:0.01)
   (@coh ?state consistency:+0.20 reason:"Methods agree")
@else
   (@coh ?state logical:-0.35 reason:"Method disagreement")
   (@refl suggestion:"Review both derivations for errors")
```

**When to use:**
- High-stakes calculations
- Novel problem types
- When domain expertise is limited

### Pattern 3: Unit Consistency Verification

**Problem:** Dimensional errors in formulas  
**Solution:** Automatic unit tracking

```lisp
;; Track units through calculation
@b ?force (@math:multiply ?mass:kg ?acceleration:"m/s²")
@assert (units ?force) == "kg⋅m/s²" == "N"

;; Coherence update on pass
(@coh ?state logical:+0.08 reason:"Dimensional analysis verified")
```

**When to use:**
- Physics/engineering calculations
- Financial modeling (currency conversions)
- Any domain with dimensional quantities

### Pattern 4: Epistemic State Recognition

**Problem:** Attempting solution when information is insufficient  
**Solution:** Goal transformation

```lisp
;; Detect hidden state without observation
@i (and (hidden_state_exists ?problem)
        (no_feedback_channel ?problem))
   
   ;; Transform goal
   @b ?original_goal "Find solution"
   @b ?new_goal "Prove no deterministic solution exists"
   
   (@refl ?state
          trigger:"epistemic_impossibility"
          suggestion:"Shift to impossibility proof")
   
   (@coh ?state 
         logical:+0.15  ; Correct problem framing
         progress:-0.20 ; But backtracking
         reason:"Goal rewrite for epistemic constraint")
```

**When to use:**
- Logic puzzles with hidden information
- Distributed systems without global view
- Underconstrained optimization problems

---

## Cross-Domain Transfer

### Example: Physics → Logic Puzzle

**Learned pattern from physics:**
```lisp
Pattern: missing_parameter_guard
Domain: physics
Trigger: completeness < 0.6 before calculation
Action: @refl "enumerate required constants"
```

**Generalized version:**
```lisp
Pattern: missing_parameter_guard_v2
Domain: any (generalized)
Trigger: completeness < 0.6 before critical_step
Action: @refl "enumerate required information"
Applicability: ["physics" "logic" "code_debugging" "planning"]
```

**Applied to logic puzzle:**
```lisp
@b ?puzzle "Three switches control three lights in separate room"
@b ?state {coh: {l:0.7 p:0.1 c:0.4}}

;; Inherited guard fires
@i (< ?state.coh.c 0.6)
   (@refl suggestion:"Enumerate information channels")
   
   @b ?available_info {
     "Switch positions": "observable",
     "Light states": "NOT observable without entering",
     "Feedback": "none until final observation"
   }
   
   ;; Recognize epistemic gap
   (@refl trigger:"no_feedback_channel"
          suggestion:"Problem may be unsolvable or require different goal")
```

### Transfer Metrics

```lisp
;; Track cross-domain effectiveness
@b ?transfer_record {
  source_domain: "physics",
  target_domain: "logic_puzzles",
  pattern: "missing_parameter_guard",
  transfer_date: "2026-01-27",
  
  effectiveness: {
    applications: 23,
    successes: 19,
    failures: 4,
    success_rate: 0.83
  },
  
  adaptations_needed: [
    "Broaden 'parameter' to 'information'",
    "Add epistemic state detection"
  ]
}
```

---

## Heuristic Lifecycle Management

### Lifecycle States

```lisp
@b ?heuristic_lifecycle {
  
  ;; Birth: Newly extracted from trace
  unverified: {
    confidence: 0.40,
    applications: 0,
    requires: "3 successful applications for promotion"
  }
  
  ;; Probation: Testing phase
  probation: {
    confidence: 0.60,
    applications: 3-9,
    promotion_threshold: "7/10 success rate",
    demotion_threshold: "3/10 fail rate"
  }
  
  ;; Promoted: Proven reliable
  promoted: {
    confidence: 0.85,
    applications: 10+,
    decay_rate: "0.01 per mismatch",
    scope: "Broad or specialized"
  }
  
  ;; Deprecated: No longer useful
  deprecated: {
    confidence: < 0.30,
    reason: ["too_narrow" "superseded" "harmful"],
    retained_for: "historical analysis"
  }
}
```

### Version Control

```lisp
;; Store heuristic with versioning
@store_heuristic ?h
  id: "missing_param_guard"
  version: "v2.1"
  confidence: 0.78
  promoted_from: "probation"
  parent_version: "v2.0"
  
  changelog: [
    "v1.0: Initial extraction from physics trace",
    "v2.0: Generalized 'parameter' → 'information'",
    "v2.1: Added epistemic state detection"
  ]
  
  decay_rules: {
    after_mismatches: 20,
    confidence_drop: 0.15
  }
  
  rollback_to: "v2.0"  ; If v2.1 fails badly
```

### Confidence Adjustment

```lisp
;; After each application
@update_heuristic ?h application_result:?result

@i (success ?result)
   (@adjust_confidence ?h delta:+0.02 cap:0.95)
@else
   (@adjust_confidence ?h delta:-0.05)
   @i (< ?h.confidence 0.50)
      (@demote ?h to:"probation")
   @i (< ?h.confidence 0.30)
      (@deprecate ?h reason:"low_effectiveness")
```

### Scope Management

```lisp
;; Narrow scope when failures occur
@b ?h_scope {
  initial: ["physics" "logic" "debugging"],
  
  after_failure_analysis: {
    removed: ["logic"],  ; Pattern doesn't work for logic puzzles
    reason: "Assumes numeric parameters, not epistemic states"
  },
  
  specialized: ["physics" "engineering" "financial_modeling"],
  
  confidence_by_domain: {
    physics: 0.92,
    engineering: 0.85,
    financial: 0.71
  }
}
```

---

## Quick Reference

### Meta-Reasoner Primitives

```lisp
;; Core operations
@meta:retrieve    ; Find similar traces
@meta:compose     ; Combine strategies
@meta:distill     ; Extract heuristics
@meta:analyze     ; Diagnose failures

;; Library management
@store_trace      ; Add to library
@tag_trace        ; Add metadata
@update_heuristic ; Adjust confidence

;; Lifecycle
@promote          ; Probation → Promoted
@demote           ; Promoted → Probation
@deprecate        ; Mark as obsolete
@rollback         ; Revert to previous version
```

### Coherence Interpretation

| Pattern | Meaning | Action |
|---------|---------|--------|
| `l ↓ p ↑` | Wrong approach | Category error - reframe problem |
| `c stalls < 0.6` | Missing info | Add gathering step or prove impossibility |
| `l high, consistency low` | Premature closure | Check edge cases |
| `All dimensions ↓` | Coherence collapse | Backtrack to last stable state |

### Retrieval Strategies

```lisp
;; Similarity-based (when domain matches)
@recall domain:?d tags:?t min_coh:0.80 sort_by:"recency"

;; Pattern-based (when task structure matches)
@recall patterns:["completeness_guard" "cross_validation"]

;; Failure-based (learn from mistakes)
@recall failure_mode:"epistemic_mismatch" include_successes:false

;; Hybrid (best of both)
@recall 
  domain:?d 
  patterns:?p 
  mix_successes:0.7 mix_failures:0.3
```

### Implementation Checklist

- [ ] **Trace storage** - JSON with full metadata
- [ ] **Coherence tracking** - Record trajectory, not just final
- [ ] **Pattern extraction** - Automated from MRL primitives
- [ ] **Retrieval system** - Vector DB + structured search
- [ ] **Failure taxonomy** - Classify errors systematically
- [ ] **Heuristic versioning** - Track confidence decay
- [ ] **Cross-domain transfer** - Test on 3+ domains
- [ ] **Meta-overhead accounting** - Prove net token savings

### Performance Targets

| Metric | Lightweight | Medium | Heavy |
|--------|-------------|---------|--------|
| Retrieval latency | <100ms | <500ms | <2s |
| Pattern extraction | Manual | Automated | Learned |
| Heuristic quality | Rule-based | Fine-tuned | RL-trained |
| Cross-domain transfer | 50% | 70% | 85% |
| Token overhead | +5% | +2% | -10% (net savings) |

---

## Advanced Topics

### Self-Modifying Meta-Reasoning

The meta-reasoner can eventually learn to modify its own retrieval and composition strategies:

```lisp
;; Meta-meta-reasoning: The system reflects on its own strategy selection
@analyze_meta_performance
  last_n_traces: 100
  metric: "coherence_gain_per_token"
  → ?meta_insights

@i (decreasing ?meta_insights.effectiveness)
   (@meta:distill focus:"retrieval_strategy"
                  suggestion:"Adjust similarity weights or add failure-based retrieval")
```

### Integration with External Systems

```lisp
;; Plugin for trace library access
(plugin
  :name "meta_reasoner"
  :command @meta
  
  :operations [
    (@meta:retrieve query:string tags:list → traces:list)
    (@meta:store trace:dict → id:string)
    (@meta:distill trace:dict → heuristic:dict)
  ]
  
  :backend {
    storage: "vector_db + structured_index"
    embedding_model: "domain-specific-retriever"
    distillation: "small-llm-or-symbolic"
  }
)
```

---

## Case Study: Three-Switch Puzzle

**Problem:** Explain why a three-switch/three-light puzzle has no deterministic solution without feedback.

### Without Meta-Reasoning (Baseline Failure)

```lisp
;; Base reasoner tries enumeration
@b ?switches ["off" "off" "off"]
@enumerate_states ?switches → ?all_combinations  ; 2³ = 8 states

;; Tries to map switches → lights
@foreach ?combo ?all_combinations
  ;; But has no way to observe lights!
  ;; Gets stuck in loop

;; Coherence profile:
{l: oscillates, p: rises then collapses, c: never exceeds 0.5}
```

**Result:** Hallucinated solution or infinite loop

### With Meta-Reasoning (Successful Failure)

```lisp
;; Meta-reasoner retrieves similar patterns
@meta:retrieve
  query: "logic puzzle information incomplete observation"
  min_final_coh: 0.60
  include_failures: true
  → ?traces

;; Extracts epistemic guard pattern
@extract ?traces pattern:"epistemic_incompleteness_guard" → ?guard

;; Injected guidance
@b ?meta_guidance {
  pattern_detected: "hidden_state_without_observation",
  forbidden_actions: ["state_enumeration_without_feedback"],
  goal_transform: "find_solution → prove_impossibility",
  expected_coherence: {
    l: "should_increase",  ; Correct framing
    p: "stays_low",        ; Not progressing to solution
    c: "moderate"          ; Have enough info about problem structure
  }
}

;; Actual reasoning with meta-guidance
@b ?problem {
  switches: 3,
  lights: 3,
  observation_channel: "single_final_check",
  feedback: "none_during_manipulation"
}

;; Apply epistemic guard
@i (and (hidden_state ?problem.lights)
        (no_feedback ?problem.observation_channel))
   
   ;; Goal rewrite
   (@refl trigger: "epistemic_impossibility"
          suggestion: "Prove no deterministic 1-to-1 mapping possible")
   
   ;; Information-theoretic argument
   @b ?analysis {
     info_needed: "log2(8) = 3 bits",  ; To distinguish 8 states
     info_available: "1 observation = log2(3) ≈ 1.58 bits",
     gap: "3 - 1.58 = 1.42 bits",
     conclusion: "Insufficient information for deterministic solution"
   }
   
   (@coh ?state
         logical: +0.40    ; Correct problem framing
         progress: 0.20    ; Not "solving" but proving
         completeness: 0.85 ; Have enough to make argument
         reason: "Transformed to information-theoretic impossibility proof")

;; Output
@o {
  conclusion: "No deterministic solution exists",
  reasoning: ?analysis,
  learned_pattern: "epistemic_guard successfully prevented enumeration trap"
}
```

### What Gets Distilled

```lisp
@meta:distill ?trace → {
  pattern: "epistemic_guard_v2",
  
  trigger: {
    hidden_state: true,
    no_observation_during_action: true,
    insufficient_bits: true
  },
  
  action: "goal_rewrite: prove_impossibility",
  
  domains: ["logic_puzzles" "distributed_systems" "incomplete_information_games"],
  
  confidence: 0.82,
  version: "v2.0"
}
```

---

## Future Directions

### Near-term (2026)

1. **Lightweight implementations** ready today
   - Vector DB for trace storage
   - Simple pattern matching
   - Rule-based heuristic extraction

2. **Medium implementations** (6-12 months)
   - Fine-tuned retrievers (ReasonIR-style)
   - Automated pattern extraction
   - Cross-domain transfer validation

### Long-term (2027+)

1. **Heavy implementations**
   - RL-based strategy learning
   - True self-improvement loops
   - Meta-reasoner writes new primitives

2. **Very heavy** (research frontier)
   - Recursive self-improvement
   - Automatic plugin generation
   - Theory of meta-reasoning

---

## License

This specification is licensed under [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/).

---

**For implementation support:**
- Repository: [github.com/mrl-lang/meta-reasoner](https://github.com/mrl-lang/meta-reasoner)
- Documentation: [docs.mrl-lang.org/meta](https://docs.mrl-lang.org/meta)
- Community: [discord.gg/mrl-meta](https://discord.gg/mrl-meta)

---

*Meta-Reasoning: Learning how to learn, one trace at a time.*
