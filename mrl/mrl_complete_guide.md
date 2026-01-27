# MRL: Complete Language Guide for LLMs
**Minimal Reasoning Language - Comprehensive Reference**  
Version 2026.1 | Last Updated: January 27, 2026

---

## Table of Contents
1. [Introduction & Philosophy](#introduction--philosophy)
2. [Core Language Specification](#core-language-specification)
3. [Plugin System](#plugin-system)
4. [Coherence-Guided Reasoning](#coherence-guided-reasoning)
5. [Practical Applications](#practical-applications)
6. [Quick Reference](#quick-reference)

---

## Introduction & Philosophy

### What is MRL?

MRL (Minimal Reasoning Language) is an ultra-compact, prefix-notation language designed specifically for LLM internal reasoning and tool orchestration. It acts as a "mental scratchpad" that bridges the gap between natural language and formal code.

### Core Design Principles

- **Token Efficiency**: 70-80% reduction compared to standard Chain-of-Thought
- **LLM-Friendly**: Prefix notation aligns with autoregressive token prediction
- **Minimal Syntax**: Small set of core primitives
- **Tool-Native**: External tools are first-class language primitives
- **Meta-Cognitive**: Built-in self-assessment and coherence tracking

### Performance Comparison

| Metric | Standard CoT | JSON-RPC | MRL |
|--------|-------------|----------|-----|
| Tokens per step | 60-100 | 80-120 | 12-20 |
| Context load | High | High | Ultra-low |
| Parsing speed | Slow | Medium | Fast |
| Logic density | Low | Low | High |
| Error detection | Reactive | Reactive | Proactive |

---

## Core Language Specification

### Essential Primitives

```lisp
@b    ; Bind - Assign values to variables
@g    ; Get - Extract data from structures
@i    ; If/When - Conditional logic
@t    ; Try - Error handling with fallback
@s    ; Synthesize - Summarize/analyze
@o    ; Output - Return results
@f    ; Fail/Assert - Error conditions
@w    ; Web search
@c    ; Code execution
```

### Basic Syntax Patterns

#### Variable Binding
```lisp
; Standard binding
@b ?data ← (@web_search "quantum computing" n:5)

; Arrow notation (alternative)
(@web_search "quantum computing" n:5 → ?data)

; Destructuring
@b {?title ?author ?year} ← (@parse_paper ?pdf)
```

#### Data Access
```lisp
; Simple field access
@g ?data:0 "title" → ?top_result

; Nested access
@g ?user "profile.settings.theme" → ?theme

; With default value
@g ?config "api_key" default:"none" → ?key
```

#### Conditional Logic
```lisp
; Basic if
@i (> ?score 0.8) (@output ?result)

; If-else
@i (exists ?cache_key)
   (@return ?cached_value)
   (@compute_fresh ?query)

; Pattern matching
@i (matches ?error "timeout") (@retry ?operation)
```

#### Error Handling
```lisp
; Try with fallback
@t (@primary_api ?query) alt:(@backup_api ?query)

; With explicit error capture
@t (@risky_operation) → ?result catch:?error

; Retry logic
(@retry 
  max_attempts:3
  delay:[1000 2000 4000]
  (@unreliable_service))
```

#### Iteration
```lisp
; Foreach
@foreach ?items
  @b ?processed (@transform ?it)
  @collect ?results ?processed

; Map pattern
@map ?items (λ [x] (@transform x))

; Filter
@filter ?items (λ [x] (> x.score 0.7))
```

### Complex Workflows

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

## Plugin System

### Plugin Definition Structure

```lisp
(plugin
  :name "weather"                 ; Unique identifier
  :command @weather              ; Invocation symbol
  :description "Get weather data for a location"
  :version "1.0"
  
  :params [
    (param :name location :type string :required true)
    (param :name days :type integer :default 1 :min 1 :max 7)
    (param :name units :type enum :options ["metric" "imperial"] :default "metric")
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

### Short Form (For Context Windows)
```lisp
; Ultra-compact plugin signature
(@weather location:string days:int=1 units:enum["metric","imperial"]=metric
  → {temp:float conditions:string forecast:list})
```

### Plugin Categories

| Prefix | Category | Examples |
|--------|----------|----------|
| `@w:*` | Web/Search | `@w:google`, `@w:arxiv` |
| `@db:*` | Database | `@db:query`, `@db:update` |
| `@api:*` | External APIs | `@api:openai`, `@api:stripe` |
| `@ml:*` | Machine Learning | `@ml:classify`, `@ml:embed` |
| `@math:*` | Mathematical | `@math:solve`, `@math:optimize` |
| `@code:*` | Code Execution | `@code:python`, `@code:js` |
| `@io:*` | Input/Output | `@io:read`, `@io:write` |
| `@util:*` | Utilities | `@util:hash`, `@util:encrypt` |

### Plugin Invocation

```lisp
; Basic usage
(@weather "London" days:3 → ?forecast)

; With options
(@web_search "quantum computing" 
  n:10
  sort:"relevance"
  domains:["arxiv.org"]
  → ?papers)

; Chained calls
(@db_query "users" where:{age: >30} → ?users)
(@analyze ?users key:"purchase_history" → ?insights)
(@format ?insights as:"report" → ?report)
```

### Type System

```lisp
; Basic types
string     "hello"
integer    42
float      3.14
boolean    true/false
list       [1 2 3]
dict       {key: "value"}
enum       ["red" "green" "blue"]

; Special types
url        "https://example.com"
email      "user@example.com"
date       "2025-01-26"
regex      "/^[A-Z]+$/"
path       "/home/user/file.txt"
```

### Error Handling Standards

```lisp
; Standard error categories
:errors {
  :validation "Invalid input parameters"
  :auth "Authentication failed"
  :rate_limit "Rate limit exceeded"
  :network "Network error"
  :timeout "Operation timed out"
}

; Error response format
(@plugin → {
  success: false
  error: "Network timeout"
  context: {
    plugin: "weather"
    params: ["London", 3]
    attempt: 2
    timestamp: "2025-01-26T10:30:00Z"
  }
})
```

### Core Plugin Library

Every MRL implementation must provide:

```lisp
; File I/O
(@io_read path:"file.txt" → ?content)
(@io_write path:"file.txt" data:?content)

; Time
(@time_now → ?timestamp)

; Utilities
(@uuid_generate → ?uuid)
(@log level:"info" message:"..." data:?context)

; Web & HTTP
(@http_get url:?url headers:?headers → ?response)
(@http_post url:?url body:?data → ?response)

; Data Processing
(@json_parse string:?json → ?data)
(@json_stringify data:?data → ?json)
(@csv_parse string:?csv → ?rows)

; Security
(@hash data:?input algorithm:"sha256" → ?hash)
(@encrypt data:?plaintext key:?key → ?ciphertext)
```

---

## Coherence-Guided Reasoning

### Meta-Cognitive Extensions

MRL includes primitives for self-assessment and reasoning quality tracking:

```lisp
@coh   ; Coherence Update - Adjust reasoning state
@refl  ; Reflection - Generate strategic pivots
@best  ; Best Selection - Choose highest coherence path
```

### Coherence State Structure

```lisp
@b ?state {
  beliefs: []                ; Current facts/assumptions
  coh: {
    l: 0.5    ; Logical consistency (0-1)
    p: 0.0    ; Progress toward goal (0-1)
    c: 0.0    ; Completeness of information (0-1)
  }
  path: "primary"           ; Current reasoning branch
  contradictions: []        ; Known conflicts
}
```

### Coherence Dimensions

| Dimension | Meaning | When to Update |
|-----------|---------|----------------|
| **Logical (l)** | Internal consistency | Evidence confirms/contradicts beliefs |
| **Progress (p)** | Movement toward goal | Getting closer to solution |
| **Completeness (c)** | Information sufficiency | New data fills gaps |

### Coherence Operations

```lisp
; Update coherence based on new evidence
@coh ?state l:+0.3 reason:"Evidence confirms hypothesis"
@coh ?state l:-0.5 reason:"Contradiction found"

; Trigger reflection when coherence is low
@i (< ?state.coh.l 0.4)
   (@refl ?state suggestion:"Revisit assumptions in step 3")

; Select best option by coherence
@best ?candidates by:"coh.l" → ?optimal_path
```

### Exploration vs. Exploitation

```lisp
; Decision logic based on coherence
@i (< ?state.coh.c 0.6)
   ; Low completeness → explore more
   (@web_search ?additional_queries n:5 → ?more_data)
   (@coh ?state c:+0.2 reason:"Added more sources")
@else
   ; High coherence → exploit current understanding
   (@s focus:"Final analysis" sources:?state.beliefs → ?conclusion)
```

---

## Practical Applications

### 1. Scientific Research & Drug Discovery

```lisp
@b ?hypothesis "Compound X inhibits Protein Y"
@b ?state {
  beliefs: ["X binds to Y in vitro", "Y overexpressed in Cancer Z"]
  coh: {l:0.7 p:0.3 c:0.4}
}

; Test prediction
@c "run_assay compound:X cell_line:Cancer_Z" → ?result

@i (> ?result.death_rate 70)
   (@coh ?state l:+0.2 p:+0.3 reason:"Assay supports hypothesis")
   (@s focus:"Strong evidence for X→Y→Cell_Death pathway")
@else
   (@coh ?state l:-0.3 reason:"Assay contradicts prediction")
   (@refl ?state suggestion:"Check binding affinity or alternative pathways")
```

**Benefits:**
- Guides exploration toward coherent mechanisms
- Flags contradictory results early
- Maintains consistency between predictions and data

### 2. Business Decision-Making

```lisp
@b ?state {
  beliefs: ["Market_A growing 15%", "Product fits need_X", "Competitor_B weak"]
  coh: {l:0.6 p:0.5 c:0.8}
  actions: ["enter_now", "partner_first", "delay_6mo"]
}

; Project outcomes for each action
@foreach ?action ?state.actions
   @b ?projection (@financial_model ?action market:"A")
   @b ?coh_gain (@estimate_coherence ?state ?projection)
   @collect ?options {action:?action gain:?coh_gain}

; Choose action maximizing coherence
@best ?options by:"gain" → ?decision
@o ?decision
```

**Benefits:**
- Balances data, projections, and strategic fit
- Surfaces incomplete/conflicting information
- Provides audit trail for decisions

### 3. Legal & Compliance Analysis

```lisp
@b ?state {
  beliefs: ["Indemnity covers all claims", "Liability capped at $10k"]
  coh: {l:0.8 p:0.1 c:0.5}
}

; Detect contradiction
@i (contradicts ?state.beliefs:0 ?state.beliefs:1)
   (@coh ?state l:-0.5 reason:"Indemnity vs cap conflict")
   (@refl ?state suggestion:"Verify precedence: Section 12.4 vs 5.2")
   @o "⚠️ RISK: Inconsistent liability scope"
```

**Benefits:**
- Automatically flags inconsistent clauses
- Maintains regulatory alignment
- Coherence score as risk indicator

### 4. Medical Diagnosis Support

```lisp
@b ?patient {
  symptoms: ["fever", "cough", "fatigue"]
  history: ["recent_travel", "no_vaccination"]
  coh: {l:0.5 p:0.6 c:0.3}
}

; Generate differential diagnoses
@b ?dx_list (@differential ?patient.symptoms)

; Score each by coherence with history
@foreach ?dx ?dx_list
   @b ?fit (@coherence_with_history ?dx ?patient.history)
   @collect ?ranked {diagnosis:?dx coherence:?fit}

; Recommend tests for low completeness
@i (< ?patient.coh.c 0.5)
   (@s suggestion:"Order CRP and chest X-ray for diagnostic clarity")
```

**Benefits:**
- Prioritizes diagnoses explaining all symptoms
- Identifies when more tests needed
- Avoids premature closure

### 5. Software Debugging

```lisp
@b ?incident "API latency spikes at 2 PM daily"
@b ?state {
  beliefs: ["CPU normal", "Memory normal", "DB queries slow"]
  coh: {l:0.4 p:0.7 c:0.6}
}

; Verify hypothesis
@t (@query_logs "slow_queries" time:"2PM" → ?queries) 
   alt:(@b ?queries null)

@i (null? ?queries)
   (@coh ?state l:-0.2 reason:"No slow queries in logs")
   (@refl suggestion:"Check network or external services")
@else
   (@coh ?state l:+0.3 p:+0.2 reason:"Found correlating queries")
   (@s focus:"Optimize queries X, Y, Z")
```

**Benefits:**
- Systematic hypothesis testing
- Tracks examined evidence
- Recognizes incoherent debug paths

### 6. Educational Tutoring

```lisp
@b ?student_work "Solve 3x + 5 = 20 → x = 10"
@b ?state {
  beliefs: ["Student subtracted 5", "Student divided by 3 incorrectly"]
  coh: {l:0.2 p:0.8 c:0.9}
}

; Identify conceptual error
@i (contradicts ?student_work "algebraic_rules")
   (@coh ?state l:-0.3 reason:"Arithmetic error in final step")
   (@refl suggestion:"Review division: (20-5)/3 = 5")
   @o "Check your division step!"
```

**Benefits:**
- Pinpoints exact reasoning breakdown
- Targeted feedback vs. generic "try again"
- Tracks conceptual progress over time

---

## Quick Reference

### When to Use MRL

✅ **Ideal For:**
- Multi-step reasoning (5+ steps)
- Tool orchestration workflows
- Token-constrained environments
- Complex research and analysis
- Mathematical/logical proofs
- Debugging and troubleshooting

❌ **Not Ideal For:**
- Simple single-step tasks
- End-user interfaces
- Human-readable documentation
- Traditional software development

### Common Patterns

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
   @b ?draft (@refine ?draft based_on:?feedback_item)
@o ?draft
```

### Performance Tips

1. **Minimize variable creation** - Reuse vars when possible
2. **Use short form plugins** - Compact signatures save tokens
3. **Chain operations** - Avoid intermediate storage
4. **Leverage coherence** - Let low scores trigger exploration
5. **Batch similar operations** - Reduce tool call overhead

### Token Budget Guidelines

| Task Complexity | MRL Tokens | Standard CoT |
|----------------|------------|--------------|
| Simple query | 10-20 | 60-100 |
| Medium analysis | 50-100 | 300-500 |
| Deep research | 200-400 | 1000-2000 |
| Full investigation | 500-1000 | 3000-5000 |

---

## Implementation Notes

### For LLM Developers

1. **Parse Incrementally**: MRL is designed for left-to-right parsing
2. **Track State**: Maintain coherence scores throughout execution
3. **Fail Fast**: Use `@f` to catch errors early
4. **Log Liberally**: Use `@log` for debugging without token cost

### For Plugin Authors

1. **Follow naming conventions** (category prefixes)
2. **Provide 3+ examples** in plugin definition
3. **Validate all inputs** before execution
4. **Return structured errors** with context
5. **Document token costs** for rate limiting

### Integration Checklist

- [ ] Plugin manifest file (`plugin.mrl`)
- [ ] Implementation code (Python/JS/etc.)
- [ ] Test suite with edge cases
- [ ] Documentation with examples
- [ ] Error handling and validation
- [ ] Rate limiting configuration
- [ ] Performance benchmarks
- [ ] Security review

---

## Version History

- **v2026.1** (Jan 2026): Initial unified specification
  - Core primitives finalized
  - Plugin system standardized
  - Coherence framework integrated
  - Cross-domain use cases documented

---

## License

This specification is licensed under [Creative Commons Attribution 4.0 International (CC BY 4.0)](https://creativecommons.org/licenses/by/4.0/).

---

**For questions, contributions, or plugin submissions:**
- Plugin Registry: [registry.mrl-lang.org](https://registry.mrl-lang.org)
- Documentation: [docs.mrl-lang.org](https://docs.mrl-lang.org)
- Community: [github.com/mrl-lang](https://github.com/mrl-lang)

---

*MRL: Think efficiently, reason coherently, act decisively.*
