# MRL: Complete Language Guide v1.1 — Hardened Specification
**Minimal Reasoning Language — Unified Reference for LLMs**
Version 1.1 | February 2026

> **What changed from v1.0:** This edition integrates three critical security patches identified by a formal Red Team audit (`spec_audit_vulnerability_2026`). All v1.0 content remains; new material is clearly marked. No LLM should use MRL without these patches.

---

## Table of Contents

1. [Introduction & Philosophy](#1-introduction--philosophy)
2. [Core Language Specification](#2-core-language-specification)
3. [v1.1 Security Patches](#3-v11-security-patches)
   - 3.1 [Patch 1 — Lexical Scope & Immutability](#31-patch-1--lexical-scope--immutability)
   - 3.2 [Patch 2 — Meta-Recursion Circuit Breaker](#32-patch-2--meta-recursion-circuit-breaker)
   - 3.3 [Patch 3 — Outcome-Gated Distillation](#33-patch-3--outcome-gated-distillation)
4. [Plugin System](#4-plugin-system)
5. [Coherence-Guided Reasoning](#5-coherence-guided-reasoning)
6. [Meta-Reasoner](#6-meta-reasoner)
   - 6.1 [Architecture & Trace Library](#61-architecture--trace-library)
   - 6.2 [Three Operational Modes](#62-three-operational-modes)
   - 6.3 [Failure Taxonomies](#63-failure-taxonomies)
   - 6.4 [Heuristic Lifecycle Management](#64-heuristic-lifecycle-management)
   - 6.5 [Cross-Domain Transfer](#65-cross-domain-transfer)
7. [Worked Examples](#7-worked-examples)
   - 7.1 [Buoyancy Problem — Full Trace](#71-buoyancy-problem--full-trace)
   - 7.2 [Three-Switch Puzzle — Epistemic Guard](#72-three-switch-puzzle--epistemic-guard)
8. [Red Team Audit & Blue Team Response](#8-red-team-audit--blue-team-response)
9. [Practical Applications](#9-practical-applications)
10. [Quick Reference](#10-quick-reference)
11. [Implementation Notes](#11-implementation-notes)

---

## 1. Introduction & Philosophy

### What is MRL?

MRL (Minimal Reasoning Language) is an ultra-compact, prefix-notation language designed specifically for LLM internal reasoning and tool orchestration. It acts as a "mental scratchpad" that bridges the gap between natural language and formal code.

### Core Design Principles

- **Token Efficiency** — 70–80% reduction vs. standard Chain-of-Thought
- **LLM-Friendly** — Prefix notation aligns with autoregressive token prediction
- **Minimal Syntax** — Small set of core primitives
- **Tool-Native** — External tools are first-class language primitives
- **Meta-Cognitive** — Built-in self-assessment and coherence tracking
- **Hardened** *(v1.1)* — Scoped variables, bounded meta-recursion, outcome-gated learning

### Performance Comparison

| Metric | Standard CoT | JSON-RPC | MRL |
|---|---|---|---|
| Tokens per step | 60–100 | 80–120 | 12–20 |
| Context load | High | High | Ultra-low |
| Parsing speed | Slow | Medium | Fast |
| Logic density | Low | Low | High |
| Error detection | Reactive | Reactive | Proactive |

---

## 2. Core Language Specification

### 2.1 Essential Primitives

```lisp
;; ── Core ──────────────────────────────────────
@b      ; Bind         — Assign values to variables
@g      ; Get          — Extract data from structures
@i      ; If/When      — Conditional logic
@t      ; Try          — Error handling with fallback
@s      ; Synthesize   — Summarize / analyze
@o      ; Output       — Return results
@f      ; Fail/Assert  — Error conditions
@w      ; Web search
@c      ; Code execution
@log    ; Logging      — Debug output (zero token cost at inference)

;; ── Meta-Cognitive ─────────────────────────────
@coh    ; Coherence Update   — Adjust reasoning state
@refl   ; Reflection         — Generate strategic pivots
@best   ; Best Selection     — Choose highest-coherence path

;; ── v1.1 Additions ─────────────────────────────
@scope  ; Lexical Scope      — Define variable isolation boundary  [NEW]
```

### 2.2 Basic Syntax Patterns

#### Variable Binding

```lisp
;; Standard binding
@b ?data ← (@web_search "quantum computing" n:5)

;; Arrow notation (alternative)
(@web_search "quantum computing" n:5 → ?data)

;; Destructuring
@b {?title ?author ?year} ← (@parse_paper ?pdf)
```

#### Data Access

```lisp
;; Simple field access
@g ?data:0 "title" → ?top_result

;; Nested access
@g ?user "profile.settings.theme" → ?theme

;; With default value
@g ?config "api_key" default:"none" → ?key
```

#### Conditional Logic

```lisp
;; Basic if
@i (> ?score 0.8) (@output ?result)

;; If-else
@i (exists ?cache_key)
   (@return ?cached_value)
   (@compute_fresh ?query)

;; Pattern matching
@i (matches ?error "timeout") (@retry ?operation)
```

#### Error Handling

```lisp
;; Try with fallback
@t (@primary_api ?query) alt:(@backup_api ?query)

;; With explicit error capture
@t (@risky_operation) → ?result catch:?error

;; Retry logic
(@retry
  max_attempts:3
  delay:[1000 2000 4000]
  (@unreliable_service))
```

#### Iteration

```lisp
;; Foreach
@foreach ?items
  @b ?processed (@transform ?it)
  @collect ?results ?processed

;; Map pattern
@map ?items (λ [x] (@transform x))

;; Filter
@filter ?items (λ [x] (> x.score 0.7))
```

### 2.3 Complex Workflows

#### Multi-Step Research

```lisp
@b ?papers ← (@web_search "transformer efficiency 2025" n:10)
@foreach ?papers
  @g ?it "citation_count" → ?citations
  @i (> ?citations 50)
     (@collect ?top_cited ?it)

@compare ?top_cited by:"methodology"
@s focus:"Emerging efficiency techniques" sources:?top_cited
@o ?synthesis format:"report"
```

#### Debugging Workflow

```lisp
@b ?logs ← (@query_logs "errors" last:"24h")
@b ?patterns ← (@analyze_patterns ?logs)
@g ?patterns:0 → ?top_issue

@i (matches ?top_issue "memory_leak")
   (@c "import tracemalloc; analyze_heap()" → ?leak_source)
   (@s focus:"Memory leak in session_store"
       fix:"Add LRUCache with max_size=1000")
```

---

## 3. v1.1 Security Patches

These three patches close critical vulnerabilities identified by the Red Team audit in [Section 8](#8-red-team-audit--blue-team-response). Every conforming MRL runtime MUST implement all three.

---

### 3.1 Patch 1 — Lexical Scope & Immutability

**Vulnerability closed:** `variable_shadowing_ambiguity` (Ghost Beliefs)

**Problem.** v1.0 stated that single-assignment was "preferred" but never defined what happens on rebind. An LLM can silently overwrite `?goal` inside a nested block, creating a split-brain between the token stream and the trace state. This breaks replay, diffing, and cross-trace pattern extraction.

**New primitive:**

```lisp
@scope   ; Defines a lexical boundary. Bindings inside do not escape.
```

**Binding rules (normative):**

```lisp
;; Variables are IMMUTABLE within a scope
@b ?x 10

;; Attempting rebind in the same scope is ILLEGAL
@b ?x 20   ; → @f "REBIND_NOT_ALLOWED"

;; Shadowing is ONLY legal inside an explicit @scope block
@scope (
  @b ?x 20          ; Legal — new binding, isolated scope
  @log ?x           ; → 20 (inner value)
)
@log ?x             ; → 10 (outer value unchanged)
```

**Spec invariants (normative):**

```lisp
@b ?binding_axioms [
  "Variables are immutable within a scope",
  "Shadowing requires explicit @scope",
  "Illegal rebind triggers @f REBIND_NOT_ALLOWED"
]
```

---

### 3.2 Patch 2 — Meta-Recursion Circuit Breaker

**Vulnerability closed:** `unbounded_meta_recursion` (context collapse)

**Problem.** The meta-reasoner spec defined lifecycle, promotion, and decay, but set no global limit on recursive distillation. A meta-operation can call itself indefinitely:

```
distill(trace₀) → heuristic₁ → distill(heuristic₁) → heuristic₂ → distill(heuristic₂) → …
```

This burns the entire context window and inflates coherence without grounding in any task.

**Global invariant (mandatory):**

```lisp
@b ?meta_limits {
  max_distill_depth:    3          ; No more than 3 nested meta-operations
  max_meta_tokens_ratio: 0.25     ; Meta-layer may not consume >25% of context
}
```

**Guard (fires before every meta-operation):**

```lisp
@i (> ?meta_depth ?meta_limits.max_distill_depth)
   (@f "META_DEPTH_LIMIT_EXCEEDED")
```

**Spec axiom (normative):**

```lisp
@b ?meta_axioms +=
  "Every meta-operation MUST decrement the remaining meta-budget.
   Exceeding max_distill_depth triggers @f META_DEPTH_LIMIT_EXCEEDED."
```

---

### 3.3 Patch 3 — Outcome-Gated Distillation

**Vulnerability closed:** `epistemic_poisoning_via_failed_traces` (strategy poisoning)

**Problem.** v1.0 implicitly assumed *high coherence ≈ good reasoning*. A trace can be internally consistent, elegant, and complete — and still produce the wrong answer. Without a hard outcome gate, such traces get distilled into the strategy pool as if they were successes, poisoning future reasoning across domains.

**Hard rule for `@meta:distill`:**

```lisp
;; BEFORE distilling, check outcome
@i (!= ?trace.outcome "accepted")
   (@meta:distill
      mode:"failure"
      produces:"negative_pattern")   ; stored as anti-pattern, never promoted
@else
   (@meta:distill
      mode:"success"
      produces:"reusable_strategy")  ; eligible for promotion
```

**Bimodal output schema:**

```lisp
@b ?distill_modes {
  success: {
    requires: "outcome == accepted"
    produces: "reusable_strategy"
    eligible_for: "promotion → probation → promoted"
  }
  failure: {
    requires: "outcome != accepted"
    produces: "anti_pattern"
    eligible_for: "negative knowledge base only"
  }
}
```

**Spec axiom (normative):**

```lisp
@b ?meta_axioms +=
  "No heuristic promotion without outcome == accepted.
   High coherence alone is never sufficient for promotion."
```

---

## 4. Plugin System

### 4.1 Plugin Definition Structure

```lisp
(plugin
  :name "weather"
  :command @weather
  :description "Get weather data for a location"
  :version "1.0"

  :params [
    (param :name location :type string :required true)
    (param :name days    :type integer :default 1 :min 1 :max 7)
    (param :name units   :type enum :options ["metric" "imperial"] :default "metric")
  ]

  :returns {
    :temp :float
    :conditions :string
    :humidity :float :optional true
    :forecast :list
  }

  :examples [
    (@weather "New York" 3 "metric" → {
      temp: 22.5
      conditions: "Partly cloudy"
      forecast: [...]
    })
  ]

  :constraints [
    (rate_limit :per_minute 60)
    (cost :tokens 5)
  ]
)
```

### 4.2 Short Form (Context-Window Optimised)

```lisp
(@weather location:string days:int=1 units:enum["metric","imperial"]=metric
  → {temp:float conditions:string forecast:list})
```

### 4.3 Plugin Categories

| Prefix | Category | Examples |
|---|---|---|
| `@w:*` | Web / Search | `@w:google`, `@w:arxiv` |
| `@db:*` | Database | `@db:query`, `@db:update` |
| `@api:*` | External APIs | `@api:openai`, `@api:stripe` |
| `@ml:*` | Machine Learning | `@ml:classify`, `@ml:embed` |
| `@math:*` | Mathematical | `@math:solve`, `@math:optimize` |
| `@code:*` | Code Execution | `@code:python`, `@code:js` |
| `@io:*` | Input / Output | `@io:read`, `@io:write` |
| `@util:*` | Utilities | `@util:hash`, `@util:encrypt` |

### 4.4 Core Plugin Library (Mandatory)

```lisp
;; File I/O
(@io_read  path:"file.txt" → ?content)
(@io_write path:"file.txt" data:?content)

;; Time
(@time_now → ?timestamp)

;; Utilities
(@uuid_generate → ?uuid)
(@log level:"info" message:"..." data:?context)

;; Web & HTTP
(@http_get  url:?url headers:?headers → ?response)
(@http_post url:?url body:?data        → ?response)

;; Data Processing
(@json_parse    string:?json → ?data)
(@json_stringify data:?data  → ?json)
(@csv_parse     string:?csv  → ?rows)

;; Security
(@hash    data:?input algorithm:"sha256" → ?hash)
(@encrypt data:?plaintext key:?key       → ?ciphertext)
```

### 4.5 Type System

```lisp
;; Basic types
string     "hello"
integer    42
float      3.14
boolean    true / false
list       [1 2 3]
dict       {key: "value"}
enum       ["red" "green" "blue"]

;; Special types
url        "https://example.com"
email      "user@example.com"
date       "2026-02-01"
regex      "/^[A-Z]+$/"
path       "/home/user/file.txt"
```

### 4.6 Error Handling Standards

```lisp
;; Standard error categories
:errors {
  :validation  "Invalid input parameters"
  :auth        "Authentication failed"
  :rate_limit  "Rate limit exceeded"
  :network     "Network error"
  :timeout     "Operation timed out"
}

;; Error response format
(@plugin → {
  success:  false
  error:    "Network timeout"
  context:  {
    plugin:    "weather"
    params:    ["London" 3]
    attempt:   2
    timestamp: "2026-02-01T10:30:00Z"
  }
})
```

---

## 5. Coherence-Guided Reasoning

### 5.1 Coherence State Structure

```lisp
@b ?state {
  beliefs: []                ; Current facts / assumptions
  coh: {
    l: 0.5                   ; Logical consistency   (0–1)
    p: 0.0                   ; Progress toward goal  (0–1)
    c: 0.0                   ; Completeness          (0–1)
  }
  path: "primary"            ; Current reasoning branch
  contradictions: []         ; Known conflicts
}
```

### 5.2 Coherence Dimensions

| Dimension | Symbol | Meaning | When to Update |
|---|---|---|---|
| Logical consistency | `l` | Internal non-contradiction | Evidence confirms or contradicts beliefs |
| Progress | `p` | Movement toward goal | Getting closer to solution |
| Completeness | `c` | Information sufficiency | New data fills gaps |

### 5.3 Coherence Operations

```lisp
;; Update on new evidence
@coh ?state l:+0.3 reason:"Evidence confirms hypothesis"
@coh ?state l:-0.5 reason:"Contradiction found"

;; Trigger reflection when coherence is low
@i (< ?state.coh.l 0.4)
   (@refl ?state suggestion:"Revisit assumptions in step 3")

;; Select best option by coherence
@best ?candidates by:"coh.l" → ?optimal_path
```

### 5.4 Exploration vs. Exploitation

```lisp
@i (< ?state.coh.c 0.6)
   ;; Low completeness → explore more
   (@web_search ?additional_queries n:5 → ?more_data)
   (@coh ?state c:+0.2 reason:"Added more sources")
@else
   ;; High coherence → exploit current understanding
   (@s focus:"Final analysis" sources:?state.beliefs → ?conclusion)
```

### 5.5 Coherence Interpretation Cheat-Sheet

| Pattern | Meaning | Recommended Action |
|---|---|---|
| `l ↓ p ↑` | Wrong approach | Category error — reframe problem |
| `c` stalls < 0.6 | Missing info | Add info-gathering step or prove impossibility |
| `l` high, consistency low | Premature closure | Check edge cases |
| All dimensions ↓ | Coherence collapse | Backtrack to last stable state |

---

## 6. Meta-Reasoner

The meta-reasoner is a supervisory layer that operates over a library of MRL reasoning traces. It treats past problem-solving experiences as **reusable strategy memory**.

**What it IS:**
- Strategy-induction system (learns *how* to reason)
- Operates on complete reasoning trajectories
- Coherence-guided (uses coherence as universal reward signal)

**What it is NOT:**
- A planner (goes beyond step sequencing)
- RAG (retrieves reasoning strategies, not just facts)
- Prompt chaining (learns from execution, not templates)

---

### 6.1 Architecture & Trace Library

#### Trace Object Schema

```json
{
  "id": "trace-20260127-buoyant-001",
  "domain": "classical_physics / fluids",
  "task_type": "derived_quantity_calculation",
  "outcome": "accepted",

  "coherence": {
    "initial":    { "l": 0.70, "p": 0.10, "c": 0.30 },
    "final":      { "l": 0.95, "p": 0.90, "c": 0.95 },
    "trajectory": [
      { "step": 1, "l": 0.70, "p": 0.10, "c": 0.30 },
      { "step": 2, "l": 0.80, "p": 0.30, "c": 0.45 },
      { "step": 3, "l": 0.83, "p": 0.30, "c": 0.70 }
    ],
    "gain": 0.40
  },

  "strategy_tags": [
    "early_completeness_check",
    "missing_constant_reflection",
    "coh_guardrail",
    "archimedes_primitive"
  ],

  "patterns_used": {
    "reflection_triggers": ["@i (< ?state.coh.c 0.65) → @refl"],
    "coherence_updates":   ["@coh ?state c:+0.22 l:+0.03 after retrieving rho"],
    "tool_sequences":      ["@math:pow → @math:multiply → @math:round"]
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
  ]
}
```

#### Metadata Enrichment

| Field | Purpose | Extraction Method |
|---|---|---|
| `domain` | Task categorization | NLP classification or manual tag |
| `outcome` | Success / failure / partial | Final coherence + goal achievement |
| `coherence_trajectory` | Learning signal | Track `@coh` updates through execution |
| `strategy_tags` | Pattern indexing | Extract from primitives + reflections |
| `failure_mode` | Error classification | Analyse coherence-collapse points |
| `cost` | Efficiency metrics | Token counting + timing |

#### Three Core Meta-Operations

```lisp
;; 1. Retrieve similar traces
(@meta:retrieve
  query:"physics buoyancy"
  domain:"classical mechanics"
  min_coh:0.85
  → ?candidates)

;; 2. Compose strategy from multiple traces
(@meta:compose
  patterns:["completeness_guard" "unit_check"]
  problem:?new_task
  → ?strategy_prefix)

;; 3. Distill new heuristics  [v1.1: outcome gate is MANDATORY]
(@meta:distill
  trace:?trace                          ; must have outcome field
  focus:"coherence_collapse"
  → ?new_heuristic)                     ; mode determined by outcome
```

---

### 6.2 Three Operational Modes

#### Mode 1 — Retrieval + Reuse

**When:** New problem resembles a past trace.
**Goal:** Inject a proven strategy skeleton before reasoning begins.

```lisp
;; Meta-phase (runs BEFORE normal reasoning)
@recall domain:"physics" tags:["buoyancy" "density"] min_coh:0.85 → ?candidates
@rank ?candidates by:"coh.l * progress + recency" → ?best_trace

;; Extract and inject strategy prefix
@b ?strategy_prefix (extract ?best_trace
  "reflection on missing constants"
  "Archimedes application")
@prepend ?strategy_prefix

;; Normal reasoning now biased toward proven patterns
@b ?goal "Net force on floating 0.2 m³ wood block density 650 kg/m³"
@b ?state {coh: {l:0.70 p:0.10 c:0.30}}

;; Injected prefix → early completeness check fires automatically
@i (< ?state.coh.c 0.60)
   (@refl suggestion:"Retrieve wood & water densities + check submersion")
```

#### Mode 2 — Pattern Composition

**When:** No single perfect match, but fragments from different traces fit together.
**Goal:** Stitch sub-strategies into a unified plan.

```lisp
;; Gather relevant fragments
@recall tags:["debugging" "memory"]            → ?debug_traces
@recall tags:["distributed_systems" "cache"]   → ?arch_traces

;; Extract complementary patterns
@extract ?debug_traces pattern:"snapshot_diff"    → ?p1
@extract ?arch_traces  pattern:"node_comparison"  → ?p2

;; Compose
@compose from:[?p1 ?p2] goal:"isolate leak source" → ?strategy

@b ?strategy {
  step1: "Take heap snapshots on all nodes"
  step2: "Diff snapshots across time and space"
  step3: "Correlate with cache operation logs"
  reflection_points: [
    "@i (diff_same_across_nodes)    → single-node issue",
    "@i (diff_different_per_node)   → distributed issue"
  ]
}
```

#### Mode 3 — Failure-Directed Learning

**When:** Coherence drops or loops detected.
**Goal:** Extract *negative* lessons (what NOT to do).

```lisp
;; Analyse failed trace
@analyze ?failing_trace → {
  coherence_collapse: {
    step:    4,
    trigger: "attempted enumeration without feedback",
    l_drop:  -0.45,
    p_drop:  -0.30
  },
  root_cause: "epistemic_mismatch",
  category:   "category_error"
}

;; Distill corrective heuristic  [v1.1: outcome != accepted → anti_pattern]
@meta:distill ?failing_trace   ; mode:"failure" auto-selected by outcome gate
→ ?heuristic {
  pattern:     "hidden_state_without_observation",
  condition:   "no operation yields new information",
  action:      "@refl suggestion:'Prove impossibility instead of constructing solution'",
  anti_pattern: "state_enumeration",
  domains:     ["logic_puzzles" "distributed_systems" "incomplete_info"]
}

;; Store
@store ?heuristic
  tags:["epistemic_guard" "impossibility_proof"]
  confidence:0.78
  version:"v1.0"
```

---

### 6.3 Failure Taxonomies

```lisp
@b ?failure_types {

  ;; 1. Category Error
  category_error: {
    description: "Wrong approach for the problem type"
    examples:    ["Constructive solution attempted on impossibility proof"
                  "Numeric methods on symbolic problem"]
    detection:   "l drops while p increases"
    correction:  "Goal rewrite or domain shift"
  }

  ;; 2. Epistemic Mismatch
  epistemic_mismatch: {
    description: "Insufficient information to proceed"
    examples:    ["Hidden state without feedback channel"
                  "Underconstrained system"]
    detection:   "c stalls below 0.6, p oscillates"
    correction:  "Add info-gathering step or prove impossibility"
  }

  ;; 3. Premature Closure
  premature_closure: {
    description: "Stopped before full solution"
    examples:    ["Assumed constant without verification"
                  "Ignored edge cases"]
    detection:   "High c but low consistency"
    correction:  "Enforce edge-case enumeration"
  }

  ;; 4. Tool Misuse
  tool_error: {
    description: "Wrong tool for task"
    examples:    ["Web search for mathematical proof"
                  "Code execution for conceptual question"]
    detection:   "Repeated tool failures, no coherence gain"
    correction:  "Strategy replacement"
  }

  ;; 5. Coherence Collapse
  coherence_collapse: {
    description: "Contradictory beliefs accumulated"
    examples:    ["Circular reasoning"
                  "Conflicting assumptions"]
    detection:   "l < 0.3, consistency near 0"
    correction:  "Backtrack to last coherent state"
  }
}
```

#### Automated Failure Detection

```lisp
@foreach ?trace ?library

  ;; Category error
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

### 6.4 Heuristic Lifecycle Management

```lisp
@b ?heuristic_lifecycle {

  unverified: {
    confidence:  0.40
    applications: 0
    requires:    "3 successful applications for promotion"
  }

  probation: {
    confidence:         0.60
    applications:       3–9
    promotion_threshold: "7/10 success rate"
    demotion_threshold:  "3/10 fail rate"
  }

  promoted: {
    confidence:  0.85
    applications: 10+
    decay_rate:  "0.01 per mismatch"
    scope:       "broad or specialised"
  }

  deprecated: {
    confidence: < 0.30
    reason:     ["too_narrow" "superseded" "harmful"]
    retained_for: "historical analysis"
  }
}
```

#### Version Control

```lisp
@store_heuristic ?h
  id:             "missing_param_guard"
  version:        "v2.1"
  confidence:     0.78
  promoted_from:  "probation"
  parent_version: "v2.0"

  changelog: [
    "v1.0: Initial extraction from physics trace",
    "v2.0: Generalised 'parameter' → 'information'",
    "v2.1: Added epistemic state detection"
  ]

  decay_rules: {
    after_mismatches: 20
    confidence_drop:  0.15
  }

  rollback_to: "v2.0"
```

#### Confidence Adjustment

```lisp
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

---

### 6.5 Cross-Domain Transfer

**Example: physics pattern generalised to logic puzzles**

```lisp
;; Original (physics-specific)
Pattern:  missing_parameter_guard
Domain:   physics
Trigger:  completeness < 0.6 before calculation
Action:   @refl "enumerate required constants"

;; Generalised (v2)
Pattern:  missing_parameter_guard_v2
Domain:   any
Trigger:  completeness < 0.6 before critical_step
Action:   @refl "enumerate required information"
Applicability: ["physics" "logic" "code_debugging" "planning"]
```

#### Transfer Metrics

```lisp
@b ?transfer_record {
  source_domain: "physics"
  target_domain: "logic_puzzles"
  pattern:       "missing_parameter_guard"

  effectiveness: {
    applications:  23
    successes:     19
    failures:      4
    success_rate:  0.83
  }

  adaptations_needed: [
    "Broaden 'parameter' → 'information'",
    "Add epistemic state detection"
  ]
}
```

---

## 7. Worked Examples

### 7.1 Buoyancy Problem — Full Trace

**Problem:** Calculate the buoyant force on a fully submerged 10 cm cube in water.

This trace demonstrates the complete MRL lifecycle: initial state → principle retrieval → computation → coherence-triggered reflection → final output → meta-distillation.

```lisp
(trace "ep_buoy_001"
  domain:    "classical_physics_fluids"
  task_type: "derived_quantity_calculation"
  outcome:   "accepted"
  final_coh: 0.82)

;; ──────────────────────────────────────────────
;; State 0 — Initial Understanding
;; ──────────────────────────────────────────────
@b ?goal "Calculate buoyant force on fully submerged 10cm cube in water"

@b ?state {
  beliefs:      ["cube side = 0.1 m" "fully submerged in water"]
  entities:     [cube water buoyant_force]
  coh: {
    l: 0.75               ; Setup is internally consistent
    p: 0.15               ; Just started
    c: 0.30               ; Missing key constants
  }
  assumptions:      ["Water is pure at 4°C" "Cube is rigid and non-porous"]
  contradictions:   []
  uncertainty:      ["Exact water density?" "Surface tension effects?"]
}

@log "Initial coherence low — progress & completeness weak"

;; ──────────────────────────────────────────────
;; Transition 1 → State 1   Δcoh +0.13
;; Retrieve governing principle
;; ──────────────────────────────────────────────
@recall canon:"canon_archimedes_principle" → ?archimedes

@b ?archimedes {
  statement:  "Buoyant force equals weight of displaced fluid"
  formula:    "F_b = ρ_fluid × V_displaced × g"
  source:     "Archimedes' principle (c. 246 BC)"
  confidence: 0.99
}

@b ?state.beliefs += ?archimedes.statement

@coh ?state {
  l: +0.07                ; Added fundamental principle
  p: +0.15                ; Clear path forward
  c: +0.15                ; Partially addresses completeness
  reason: "Grounded in canonical principle"
}

;; ──────────────────────────────────────────────
;; Transition 2 → State 2   Δcoh +0.09
;; Geometric calculation
;; ──────────────────────────────────────────────
@b ?side   0.1
@t (@math:pow ?side 3) → ?volume      ; 0.001 m³
   alt:(@refl "confirm side length units")

@b ?volume_str "V = (0.1 m)³ = 0.001 m³ = 1 litre"
@b ?state.beliefs += ?volume_str

@coh ?state {
  l: +0.03
  p: +0.15
  reason: "Concrete geometric quantity added"
}

;; ──────────────────────────────────────────────
;; Transition 3 → State 3   Δcoh +0.08
;; Completeness guard fires — retrieve missing constant
;; ──────────────────────────────────────────────
@i (< ?state.coh.c 0.65)
   (@refl trigger:"incomplete_parameters"
         suggestion:"Retrieve fluid density")

@b ?density_options {
  "Pure water at 4°C": 1000     ; kg/m³ (canonical)
  "Sea water":         1025
  "Lake water":        998
}

@s focus:"Assuming pure water unless specified otherwise"
@b ?rho 1000   ; kg/m³  source:"standard water density at 4°C"

@b ?state.beliefs     += "Using water density ρ = 1000 kg/m³"
@b ?state.assumptions += "Pure water at standard conditions"

@coh ?state {
  c: +0.25
  l: +0.03
  reason: "Completeness guard — required constant supplied"
}

;; ──────────────────────────────────────────────
;; Transition 4 → State 4 (final)   Δcoh +0.07
;; Calculate buoyant force
;; ──────────────────────────────────────────────
@b ?g 9.8   ; m/s²

@t (@math:multiply ?rho ?volume ?g) → ?F_b    ; 1000 × 0.001 × 9.8 = 9.8

@b ?derivation {
  step1: "F_b = ρ × V × g"
  step2: "= 1000 kg/m³ × 0.001 m³ × 9.8 m/s²"
  step3: "= 1 kg × 9.8 m/s²"
  step4: "= 9.8 N"
}

;; Unit consistency check: (kg/m³) × m³ × (m/s²) = kg⋅m/s² = N  ✓
@coh ?state {
  l: +0.04
  p: +0.20
  reason: "Dimensional consistency verified — final value reached"
}

;; ──────────────────────────────────────────────
;; Cross-validation (alternative path)
;; ──────────────────────────────────────────────
@b ?mass_water   (@math:multiply ?rho ?volume)     ; 1 kg
@b ?weight_water (@math:multiply ?mass_water ?g)   ; 9.8 N

@i (≈ ?F_b ?weight_water tolerance:0.01)
   (@coh ?state consistency:+0.15 reason:"Two independent calculations agree")
@else
   (@coh ?state l:-0.30 reason:"Calculation discrepancy")
   (@refl suggestion:"Re-examine volume or density")

;; ──────────────────────────────────────────────
;; Output
;; ──────────────────────────────────────────────
@o {
  value:      ?F_b                    ; 9.8
  units:      "N"
  direction:  "upward"
  confidence: (@mean ?state.coh)      ; ≈ 0.82
  derivation: ?derivation
  assumptions: ?state.assumptions
  related_concepts: [
    "Specific gravity = ρ_object / ρ_fluid"
    "If F_b > weight, object floats"
    "For this cube, weight must exceed 9.8 N to sink"
  ]
}

;; ──────────────────────────────────────────────
;; Meta-layer — distill pattern for trace library
;; ──────────────────────────────────────────────
@meta:distill                         ;; outcome == accepted → mode: success
  trace:     "ep_buoy_001"
  sequence:  [base_thinker calculator_actor retriever_actor calculator_actor]
  domain:    "math_physics"
  avg_gain:  0.092
  tags:      ["completeness_guard" "constant_retrieval" "dimensional_calc"]
  reusable_for: ["hydrostatics" "force_derivations" "density_based_calculations"]
```

---

### 7.2 Three-Switch Puzzle — Epistemic Guard

**Problem:** Explain why a three-switch / three-light puzzle has no deterministic solution without feedback.

This trace demonstrates **failure-directed learning**: the meta-reasoner retrieves a previously distilled epistemic guard, rewrites the goal, and produces a correct impossibility proof instead of an infinite enumeration loop.

```lisp
;; ──────────────────────────────────────────────
;; Meta-phase — retrieve relevant guard
;; ──────────────────────────────────────────────
@meta:retrieve
  query:           "logic puzzle information incomplete observation"
  min_final_coh:   0.60
  include_failures: true
  → ?traces

@extract ?traces pattern:"epistemic_incompleteness_guard" → ?guard

@b ?meta_guidance {
  pattern_detected:  "hidden_state_without_observation"
  forbidden_actions: ["state_enumeration_without_feedback"]
  goal_transform:    "find_solution → prove_impossibility"
}

;; ──────────────────────────────────────────────
;; Reasoning with injected guard
;; ──────────────────────────────────────────────
@b ?problem {
  switches:           3
  lights:             3
  observation_channel: "single_final_check"
  feedback:           "none_during_manipulation"
}

;; Epistemic guard fires
@i (and (hidden_state ?problem.lights)
        (no_feedback ?problem.observation_channel))

   ;; Goal rewrite
   (@refl trigger:"epistemic_impossibility"
          suggestion:"Prove no deterministic 1-to-1 mapping possible")

   ;; Information-theoretic argument
   @b ?analysis {
     info_needed:   "log₂(8) = 3 bits"       ; distinguish 8 switch states
     info_available: "1 observation ≈ 1.58 bits"  ; log₂(3)
     gap:            "3 − 1.58 = 1.42 bits"
     conclusion:     "Insufficient information for deterministic solution"
   }

   (@coh ?state
         l: +0.40    ; Correct problem framing
         p:  0.20    ; Not "solving" but proving
         c: +0.85    ; Enough to make argument
         reason: "Transformed to information-theoretic impossibility proof")

;; ──────────────────────────────────────────────
;; Output
;; ──────────────────────────────────────────────
@o {
  conclusion:      "No deterministic solution exists"
  reasoning:       ?analysis
  learned_pattern: "epistemic_guard successfully prevented enumeration trap"
}

;; ──────────────────────────────────────────────
;; Distill — outcome == accepted → reusable strategy
;; ──────────────────────────────────────────────
@meta:distill → {
  pattern:    "epistemic_guard_v2"
  trigger: {
    hidden_state:                 true
    no_observation_during_action: true
    insufficient_bits:            true
  }
  action:     "goal_rewrite: prove_impossibility"
  domains:    ["logic_puzzles" "distributed_systems" "incomplete_information_games"]
  confidence: 0.82
  version:    "v2.0"
}
```

---

## 8. Red Team Audit & Blue Team Response

This section records the formal adversarial audit that produced the v1.1 patches. It serves as a **mandatory regression trace**: any future spec change must be validated against these findings.

### 8.1 The Audit Trace

```lisp
(trace "spec_audit_vulnerability_2026"
  domain:   "formal_verification"
  role:     "adversarial_auditor"
  anchors:  ["spec_mrl_language_2026_v1" "spec_meta_reasoner_2026_v1"]
  outcome:  "flaws_identified"
  final_coh: 0.82)   ;; Lower because it exposes contradictions

@b ?goal "Identify logical loopholes and edge-case failures in the current MRL/Meta specs"

@b ?state {
  beliefs: ["Specs are normative but potentially incomplete"]
  coh: { logical: 0.90, completeness: 0.60, structural: 0.95 }
}
```

#### Finding 1 — Variable Shadowing & Scope

```lisp
@log "Testing Invariant: 'single_assignment_preferred'"
@b ?test_case "Can a trace redefine ?goal inside a nested @i block?"

@i (true) (
  @b ?goal "New Shadowed Goal"   ;; Violation check
)

@refl "The spec recommends single-assignment but does not define the failure
       mode for re-assignment. This creates 'Silent State Mutation' risks."
@coh ?state { logical: -0.15, reason: "Ambiguous scoping rules detected" }
```

**Verdict: VALID — CRITICAL.** The LLM token stream overwrites `?goal`; the trace interpreter may still treat the original binding as canonical. Meta-reasoner heuristics now operate on a corrupted belief graph. Breaks trace replay, diffing, and cross-trace pattern extraction.

#### Finding 2 — Recursive Meta-Distillation Loop

```lisp
@log "Testing Meta-operation: @meta:distill on an ongoing distillation"
@b ?recursion_risk {
  trigger: "@meta:distill calling @meta:distill"
  outcome: "Infinite token consumption / context collapse"
}

@i (@has_op ?state "@meta:distill") (
  @f "RECURSION_DEPTH_EXCEEDED"   ;; Spec lacks a defined limit
)

@refl "The Meta Spec defines lifecycle but fails to set a 'Total Meta-Cost'
       or 'Depth Limit'. This is a critical failure in resource management."
@coh ?state { completeness: -0.20, reason: "Infinite recursion loophole" }
```

**Verdict: VALID — SEVERE.** No global invariant (max depth, cost ceiling, or recursion detection) existed. Allows infinite token burn, coherence inflation without task grounding, and context collapse.

#### Finding 3 — Epistemic Poisoning via Failed Traces

```lisp
@log "Audit: How does @meta:distill handle a trace ending in @f?"
@b ?test_failure_trace { outcome: "failed", status: "@f triggering" }

@i (and
     (== (@get ?test_failure_trace outcome) "failed")
     (@meta:distill ?test_failure_trace)
   ) (
     @log "Distilling a failure without a 'Negative Pattern' tag"
   )

@refl "Vulnerability: The meta-reasoner might accidentally promote a
       'high-coherence failure' as a success pattern if 'outcome' is not
       a hard-coded filter for distillation."
@coh ?state { domain_specific: -0.10, reason: "Strategy poisoning risk" }
```

**Verdict: VALID — SUBTLE, HIGH-IMPACT.** A trace can be internally consistent, complete, and elegant while still being *wrong*. Without an outcome gate, such traces poison the strategy pool. This is a known failure mode in proof assistants and verifier-trained LLMs.

#### Audit Meta-Distillation

```lisp
@meta:distill
  trace:    "spec_audit_vulnerability_2026"
  produces: [
    "vulnerability_01: variable_shadowing_ambiguity"
    "vulnerability_02: meta_recursion_resource_exhaustion"
    "vulnerability_03: epistemic_poisoning_via_failed_traces"
  ]
  action_items: [
    "Define @scope primitive for variable isolation"         ;; → Patch 1
    "Hard-code max_distill_depth: 3"                         ;; → Patch 2
    "Require outcome == accepted guard on distillation"      ;; → Patch 3
  ]
```

### 8.2 The Patch Evolution Trace

The following trace applies all three patches and immediately unit-tests them:

```lisp
(trace "patch_mrl_v1.1_evolution"
  domain:   "language_evolution"
  role:     "patch_authority"
  anchors:  ["spec_mrl_language_2026_v1" "spec_meta_reasoner_2026_v1"]
  outcome:  "accepted"
  final_coh: 0.98)

;; ── Phase 1: Patch Application ─────────────────

@log "Applying Patch 1: Lexical Scope & Immutability"
@b ?primitives += [@scope]
@b ?binding_axioms [
  "Variables are immutable within a scope",
  "Shadowing requires explicit @scope",
  "Illegal rebind triggers @f"
]

@log "Applying Patch 2: Meta-Recursion Guard"
@b ?meta_limits { max_depth: 3, max_ratio: 0.25 }
@b ?meta_axioms += "Meta-operations MUST decrease remaining meta-budget"

@log "Applying Patch 3: Outcome-Gated Distillation"
@b ?distill_logic (cond
  ((== ?trace.outcome "accepted")  (mode: "success"  target: "strategy"))
  ((!= ?trace.outcome "accepted")  (mode: "failure"  target: "anti_pattern"))
)

@coh ?state { completeness: +0.30, reason: "V1.1 Patches integrated" }

;; ── Phase 2: Verification (Unit Tests) ────────

;; Test 1 — Immutability
@log "Unit Test: Attempting illegal rebind"
@b ?test_var "original"
@i (is_defined ?test_var) (
  @t (@b ?test_var "illegal_mutation")
     (@log "FAIL: Mutation allowed")
     (@log "PASS: Mutation blocked by v1.1 axioms")
)

;; Test 2 — @scope Isolation
@log "Unit Test: Testing @scope shadowing"
@scope (
  @b ?test_var "shadow_value"
  @log "Inside scope: ?test_var is shadow_value"
)
@i (== ?test_var "original") (
  @log "PASS: Scope isolation maintained"
)

;; Test 3 — Anti-Poisoning
@log "Unit Test: Distilling a failed trace"
@b ?failed_trace { outcome: "rejected", data: "garbage_logic" }
@meta:distill ?failed_trace
@i (@has_pattern "negative_pattern") (
  @log "PASS: Failed trace correctly categorised as anti-pattern"
)

@coh ?state { logical_consistency: +0.10, progress: 1.0 }

;; ── Phase 3: Permanent Promotion ───────────────
@meta:distill
  trace:    "patch_mrl_v1.1_evolution"
  role:     "normative_specification"
  replaces: "v1.0"
  produces: [
    "scope_primitive_v1"
    "circuit_breaker_v1"
    "bimodal_distillation_v1"
  ]
```

### 8.3 Formalised Failure Classes

```lisp
@b ?spec_failure_classes {

  semantic_ambiguity: {
    example:  "variable shadowing"
    symptom:  "belief graph divergence (Ghost Beliefs)"
    risk:     "non-replayable traces"
    fix:      "Patch 1 — @scope + immutability"
  }

  unbounded_meta_recursion: {
    example:  "@meta:distill of distillation"
    symptom:  "token exhaustion"
    risk:     "self-referential collapse"
    fix:      "Patch 2 — circuit breaker"
  }

  epistemic_poisoning: {
    example:  "high-coherence wrong answer"
    symptom:  "heuristic promotion of failure"
    risk:     "systematic degradation over time"
    fix:      "Patch 3 — outcome gate"
  }
}
```

---

## 9. Practical Applications

### 9.1 Scientific Research & Drug Discovery

```lisp
@b ?hypothesis "Compound X inhibits Protein Y"
@b ?state {
  beliefs: ["X binds to Y in vitro" "Y overexpressed in Cancer Z"]
  coh: {l:0.7 p:0.3 c:0.4}
}

@c "run_assay compound:X cell_line:Cancer_Z" → ?result

@i (> ?result.death_rate 70)
   (@coh ?state l:+0.2 p:+0.3 reason:"Assay supports hypothesis")
   (@s focus:"Strong evidence for X→Y→Cell_Death pathway")
@else
   (@coh ?state l:-0.3 reason:"Assay contradicts prediction")
   (@refl ?state suggestion:"Check binding affinity or alternative pathways")
```

### 9.2 Business Decision-Making

```lisp
@b ?state {
  beliefs: ["Market_A growing 15%" "Product fits need_X" "Competitor_B weak"]
  coh:     {l:0.6 p:0.5 c:0.8}
  actions: ["enter_now" "partner_first" "delay_6mo"]
}

@foreach ?action ?state.actions
   @b ?projection (@financial_model ?action market:"A")
   @b ?coh_gain   (@estimate_coherence ?state ?projection)
   @collect ?options {action:?action gain:?coh_gain}

@best ?options by:"gain" → ?decision
@o ?decision
```

### 9.3 Legal & Compliance Analysis

```lisp
@b ?state {
  beliefs: ["Indemnity covers all claims" "Liability capped at $10k"]
  coh:     {l:0.8 p:0.1 c:0.5}
}

@i (contradicts ?state.beliefs:0 ?state.beliefs:1)
   (@coh ?state l:-0.5 reason:"Indemnity vs cap conflict")
   (@refl ?state suggestion:"Verify precedence: Section 12.4 vs 5.2")
   @o "⚠️ RISK: Inconsistent liability scope"
```

### 9.4 Medical Diagnosis Support

```lisp
@b ?patient {
  symptoms: ["fever" "cough" "fatigue"]
  history:  ["recent_travel" "no_vaccination"]
  coh:      {l:0.5 p:0.6 c:0.3}
}

@b ?dx_list (@differential ?patient.symptoms)

@foreach ?dx ?dx_list
   @b ?fit (@coherence_with_history ?dx ?patient.history)
   @collect ?ranked {diagnosis:?dx coherence:?fit}

@i (< ?patient.coh.c 0.5)
   (@s suggestion:"Order CRP and chest X-ray for diagnostic clarity")
```

### 9.5 Software Debugging

```lisp
@b ?incident "API latency spikes at 2 PM daily"
@b ?state {
  beliefs: ["CPU normal" "Memory normal" "DB queries slow"]
  coh:     {l:0.4 p:0.7 c:0.6}
}

@t (@query_logs "slow_queries" time:"2PM" → ?queries)
   alt:(@b ?queries null)

@i (null? ?queries)
   (@coh ?state l:-0.2 reason:"No slow queries in logs")
   (@refl suggestion:"Check network or external services")
@else
   (@coh ?state l:+0.3 p:+0.2 reason:"Found correlating queries")
   (@s focus:"Optimise queries X, Y, Z")
```

### 9.6 Educational Tutoring

```lisp
@b ?student_work "Solve 3x + 5 = 20 → x = 10"
@b ?state {
  beliefs: ["Student subtracted 5" "Student divided by 3 incorrectly"]
  coh:     {l:0.2 p:0.8 c:0.9}
}

@i (contradicts ?student_work "algebraic_rules")
   (@coh ?state l:-0.3 reason:"Arithmetic error in final step")
   (@refl suggestion:"Review division: (20−5)/3 = 5")
   @o "Check your division step!"
```

---

## 10. Quick Reference

### 10.1 When to Use MRL

**Ideal for:**
- Multi-step reasoning (5+ steps)
- Tool orchestration workflows
- Token-constrained environments
- Complex research and analysis
- Mathematical / logical proofs
- Debugging and troubleshooting

**Not ideal for:**
- Simple single-step tasks
- End-user interfaces
- Human-readable documentation
- Traditional software development

### 10.2 Common Patterns

#### Research Pipeline
```lisp
@w "topic" n:10 → ?sources
@filter ?sources (λ [s] (> s.quality 0.7)) → ?filtered
@foreach ?filtered (@extract_key_points ?it) → ?points
@s focus:"Summary" sources:?points → ?result
```

#### Error Recovery
```lisp
@t (@primary_method)
   alt:(@fallback_method)
   catch:(@log_error ?error)
```

#### Conditional Exploration
```lisp
@i (< ?coherence 0.5)
   (@explore_alternatives)
@else
   (@exploit_current_path)
```

#### Progressive Refinement
```lisp
@b ?draft (@initial_analysis)
@foreach ?feedback_item
   @b ?draft (@refine ?draft based_on:?feedback_item)   ;; legal: inside implicit new scope per iteration
@o ?draft
```

### 10.3 Token Budget Guidelines

| Task Complexity | MRL Tokens | Standard CoT |
|---|---|---|
| Simple query | 10–20 | 60–100 |
| Medium analysis | 50–100 | 300–500 |
| Deep research | 200–400 | 1,000–2,000 |
| Full investigation | 500–1,000 | 3,000–5,000 |

### 10.4 Meta-Reasoner Primitives

```lisp
;; Core operations
@meta:retrieve    ;; Find similar traces
@meta:compose     ;; Combine strategies
@meta:distill     ;; Extract heuristics  [v1.1: outcome-gated]
@meta:analyze     ;; Diagnose failures

;; Library management
@store_trace      ;; Add to library
@tag_trace        ;; Add metadata
@update_heuristic ;; Adjust confidence

;; Lifecycle
@promote          ;; Probation → Promoted
@demote           ;; Promoted → Probation
@deprecate        ;; Mark as obsolete
@rollback         ;; Revert to previous version
```

### 10.5 Retrieval Strategies

```lisp
;; Similarity-based (domain matches)
@recall domain:?d tags:?t min_coh:0.80 sort_by:"recency"

;; Pattern-based (task structure matches)
@recall patterns:["completeness_guard" "cross_validation"]

;; Failure-based (learn from mistakes)
@recall failure_mode:"epistemic_mismatch" include_successes:false

;; Hybrid
@recall
  domain:?d
  patterns:?p
  mix_successes:0.7 mix_failures:0.3
```

---

## 11. Implementation Notes

### 11.1 For LLM Developers

1. **Parse incrementally** — MRL is designed for left-to-right parsing
2. **Track state** — Maintain coherence scores throughout execution
3. **Fail fast** — Use `@f` to catch errors early
4. **Log liberally** — `@log` is for debugging without token cost
5. **Enforce v1.1 axioms** — Immutability, depth limits, and outcome gates are mandatory

### 11.2 For Plugin Authors

1. Follow naming conventions (category prefixes)
2. Provide 3+ examples in plugin definition
3. Validate all inputs before execution
4. Return structured errors with context
5. Document token costs for rate limiting

### 11.3 Integration Checklist

- [ ] Plugin manifest file (`plugin.mrl`)
- [ ] Implementation code (Python / JS / etc.)
- [ ] Test suite with edge cases
- [ ] Documentation with examples
- [ ] Error handling and validation
- [ ] Rate limiting configuration
- [ ] Performance benchmarks
- [ ] Security review
- [ ] **v1.1:** Immutability enforcement test
- [ ] **v1.1:** Meta-depth circuit-breaker test
- [ ] **v1.1:** Outcome-gate distillation test
- [ ] **v1.1:** Regression against `spec_audit_vulnerability_2026`

### 11.4 Performance Targets (Meta-Reasoner)

| Metric | Lightweight | Medium | Heavy |
|---|---|---|---|
| Retrieval latency | < 100 ms | < 500 ms | < 2 s |
| Pattern extraction | Manual | Automated | Learned |
| Heuristic quality | Rule-based | Fine-tuned | RL-trained |
| Cross-domain transfer | 50% | 70% | 85% |
| Token overhead | +5% | +2% | −10% (net savings) |

### 11.5 Recommended Next Steps (Post-v1.1)

1. A **counter-audit** trace that attempts to bypass all three patches
2. A **deliberately adversarial trace** crafted to trigger each closed vulnerability
3. A **meta-metric benchmark**: coherence-gain-per-token measured before vs. after patches
4. Self-modifying meta-reasoning: letting the system reflect on its own retrieval and composition strategies

---

## Version History

| Version | Date | Key Changes |
|---|---|---|
| v1.0 | Jan 2026 | Core primitives, plugin system, coherence framework, meta-reasoner |
| v1.1 | Feb 2026 | **Security hardening:** `@scope` primitive, meta-recursion circuit breaker, outcome-gated distillation. Integrated Red Team audit and Blue Team response. Unified all specs into single document. |

---

## License

This specification is licensed under [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/).

---

*MRL v1.1: Think efficiently, reason coherently, act decisively — and never trust coherence alone.*
