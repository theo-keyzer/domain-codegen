# Runtime Insights: Current Mechanics → Generalized Principles

## Core Mechanism 1: The Win Stack as Consciousness

### How It Currently Works

```go
type WinT struct {
    Name     string      // Actor being matched
    Dat      interface{} // Current data context
    Cnt      int         // How many times actor fired
    Arg      string      // Passed argument
    IsPrev   bool        // Previous actor succeeded
    IsOn     bool        // Delay mode active
    IsTrig   bool        // Trigger fired
}
```

The Win stack is a **linear array of execution contexts** where:
- Each frame holds a moment of "attention" (Name + Dat)
- Frames can reach backward via `${.actor_name.field}` 
- Triggers can replay earlier frames

### Generalized Principle

**Context Stack Pattern**: Any code generation system needs layered contexts where:

| Layer | Contains | Access Pattern |
|-------|----------|----------------|
| Current | Immediate data | `${field}` |
| Parent | Enclosing scope | `${parent.field}` |
| Named | Any prior context | `${.name.field}` |
| Global | Shared state | `_.collection` |

```
┌─────────────────────────────────────┐
│ Win Stack = Addressable Memory      │
│                                     │
│  winp=3: Actor "gen_field"          │
│          Dat = current Element      │
│          ${name} → this element     │
│                    ↑                │
│  winp=2: Actor "gen_comp"           │
│          Dat = parent Comp          │
│          ${.gen_comp.name} ─────────┘
│                    ↑                │
│  winp=1: Actor "main"               │
│          Dat = root                 │
│                    ↑                │
│  winp=0: Entry point                │
└─────────────────────────────────────┘
```

**Could Work Better**: Make backward references explicit and type-safe:
```
${@parent.name}      // Structural parent (Kparentp)
${@context.name}     // Win stack walk (current behavior)
${@caller.name}      // Immediate caller only
${@root.name}        // Always winp=0
```

---

## Core Mechanism 2: Conditional Actor Matching

### How It Currently Works

```go
for i := 0; i < len(glob.Acts.ApActor); i++ {
    if glob.Acts.ApActor[i].Kname != name {
        continue
    }
    
    act := glob.Acts.ApActor[i]
    if act.Kattr != "E_O_L" {
        // Evaluate condition
        valOk, val := strs(glob, winp, act.Kvalue, ...)
        varOk, varv := sGetVar(glob, winp, ...)
        
        if !chk(glob, act.Keq, varv, val, prev, varOk, valOk, ...) {
            prev = false
            continue
        }
    }
    
    prev = true
    // Execute matching actor's commands
}
```

**Key Insight**: All actors with matching name are tried in order. First match doesn't stop iteration—**all matching actors execute**.

### Generalized Principle

**Multi-Dispatch Pattern**: Instead of single-dispatch OOP, use pattern-matching dispatch:

```
┌─────────────────────────────────────────────────┐
│ Actor Name = Dispatch Key                       │
│ Condition  = Guard Clause                       │
│ Body       = Action                             │
│                                                 │
│ All guards evaluated, all matching bodies run   │
└─────────────────────────────────────────────────┘
```

| Traditional OOP | Actor System |
|-----------------|--------------|
| One method per signature | Multiple actors per name |
| First match wins | All matches execute |
| Explicit if/switch | Declarative conditions |
| Override replaces | Chains compose |

**Could Work Better**: Add explicit dispatch modes:
```
Actor gen_field Element type = ref        # Mode: all (default)
Actor.first gen_special Element name = id # Mode: first match only
Actor.last gen_fallback Element          # Mode: only if no others matched
```

---

## Core Mechanism 3: The Trigger/Delay System

### How It Currently Works

```go
// Out delay sets up delayed output
if kwhat == "delay" {
    glob.Wins[winp].IsOn = true
    glob.Wins[winp].OnPos = i  // Remember command position
}

// C command checks if should trigger
case *KpC:
    if glob.Wins[winp].IsOn && !glob.Wins[winp].IsTrig {
        continue  // Skip if delayed and not triggered
    }
    trig(glob, winp)  // Trigger parent chain

// trig() replays delayed commands
func trig(glob *GlobT, winp int) {
    if !glob.Wins[winp].IsPrev || winp == 0 {
        return
    }
    glob.Wins[prev].IsTrig = true
    reGoCmds(glob, prev)  // Replay from OnPos to CurPos
}
```

**Key Insight**: This is **lazy evaluation with memoization**. Output is deferred until we know it's needed.

### Generalized Principle

**Conditional Section Emission**:

```
┌─────────────────────────────────────────────────────┐
│ Problem: Don't emit "Section Header" if no content │
│                                                     │
│ Solution: Delay header, trigger when content found │
│                                                     │
│ Out delay                                           │
│ C Section: ${name}    ← Recorded but not emitted   │
│ Its Element gen_elem  ← If any match...            │
│                                                     │
│ Actor gen_elem        ← First C here triggers      │
│   C   Field: ${name}  ← This triggers parent's C   │
└─────────────────────────────────────────────────────┘
```

**Could Work Better**: Make the pattern explicit and nestable:

```
Section ${name}              # Implicit delay
  Its Element gen_elem       # Content determines emission
EndSection                   # Explicit boundary

# Or functional style:
IfContent
  C Header
  Its Element gen_body
  C Footer
EndIf
```

---

## Core Mechanism 4: Path Resolution

### How It Currently Works

```go
func sGetVar(glob *GlobT, winp int, sc []string, va []string, lno string) (bool, string) {
    path := va
    rec := getPath(glob, winp, path, lno)  // Walk the stack
    
    if kp, ok := dat.(Kp); ok && len(rec.Path) > 0 {
        resOk, res := kp.GetVar(glob, rec.Path, lno)  // Delegate to component
        dat = res
    }
    
    return cmdVar(glob, sc, dat, 3)  // Apply modifiers (:c, :l, :u)
}
```

**Key Insight**: Resolution is two-phase:
1. `getPath`: Walk Win stack to find context
2. `GetVar`: Walk component structure within context

### Generalized Principle

**Layered Resolution**:

```
${.gen_comp.parent.elements.0.name:c}
  │         │      │        │ │    └─ Modifier: capitalize
  │         │      │        │ └────── Field: name
  │         │      │        └──────── Index: first element
  │         │      └───────────────── Navigate: elements array
  │         └──────────────────────── Navigate: parent pointer
  │         └──────────────────────── Context: find gen_comp in stack
  └────────────────────────────────── Prefix: stack walk mode
```

| Prefix | Resolution |
|--------|------------|
| (none) | Current context, then walk up |
| `.name` | Find named actor in stack |
| `_` | Global collection |
| `@` | Could be: explicit mode |

**Could Work Better**: Unify all navigation under one grammar:

```
${path}                    # Current context
${^path}                   # Parent context
${^^path}                  # Grandparent
${@actor.path}             # Named context
${_.collection.key}        # Global collection
${#Comp.path}              # Type-based lookup
```

---

## Core Mechanism 5: The `chk` Function (Condition Evaluation)

### How It Currently Works

```go
func chk(glob *GlobT, eqa string, v, ss interface{}, prev, attrOk, valOk bool, lno string) bool {
    eq := eqa
    
    // Existence checks
    if eq == "??" { return !attrOk || !valOk }  // Missing
    if eq[0] == '?' { if !attrOk || !valOk { return false }; eq = eq[1:] }  // Exists
    
    // Chain operators
    if eq[0] == '&' { if !prev { return false }; eq = eq[1:] }  // AND previous
    if eq[0] == '|' { if prev { return true }; eq = eq[1:] }    // OR previous
    
    // Comparison operators
    switch eq {
    case "=":    return vStr == ssStr
    case "!=":   return vStr != ssStr
    case "in":   return contains(ssStr, vStr)
    case "!in":  return !contains(ssStr, vStr)
    case "has":  return contains(vStr, ssStr)
    case "regex": return regexp.Match(ssStr, vStr)
    }
}
```

**Key Insight**: Operators are composable prefixes: `?&=` means "exists AND previous AND equals"

### Generalized Principle

**Composable Condition Algebra**:

```
┌─────────────────────────────────────────────────────────┐
│ Prefix Layer (modify behavior):                         │
│   ?  = require attribute exists                         │
│   ?? = require attribute missing                        │
│   &  = require previous actor matched                   │
│   |  = short-circuit if previous matched                │
│                                                         │
│ Operator Layer (compare values):                        │
│   =, !=, <, >, <=, >=                                   │
│   in, !in, has, !has                                    │
│   regex, glob, starts, ends                             │
│                                                         │
│ Composition: prefix* operator                           │
│   ?&=   → exists AND previous AND equals                │
│   |!=   → OR-with-previous, if checking then not-equals │
└─────────────────────────────────────────────────────────┘
```

**Could Work Better**: Make the grammar explicit and extensible:

```
Actor x Element 
  when exists(name) 
  and prev 
  and type = ref

# Or keep compact but document clearly:
Actor x Element name ? type &= ref
#                    │      │└─ equals "ref"
#                    │      └── AND previous matched  
#                    └───────── name must exist
```

---

## Core Mechanism 6: Du (Do/Invoke) Inheritance

### How It Currently Works

```go
case *KpDu:
    NewAct(glob, c.Kactor, args, c.LineNo)
    // Copy context from current frame
    glob.Wins[winp+1].Cnt = glob.Wins[winp].Cnt
    glob.Wins[winp+1].PrevCnt = glob.Wins[winp].PrevCnt
    glob.Wins[winp+1].DataKey = glob.Wins[winp].DataKey
    glob.Wins[winp+1].DataType = glob.Wins[winp].DataType
    glob.Wins[winp+1].DataKeys = glob.Wins[winp].DataKeys
    ret := GoAct(glob, glob.Wins[winp].Dat)  // Same Dat!
```

**Key Insight**: `Du` creates a new actor frame but **shares the same data context**. It's like calling a subroutine that operates on the same component.

### Generalized Principle

**Delegation Modes**:

| Command | Data Context | Win Frame | Use Case |
|---------|--------------|-----------|----------|
| `Du actor` | Same | New (inherits Cnt) | Subroutine call |
| `Its path actor` | Children | New (fresh) | Iteration |
| `All Type actor` | All of type | New (fresh) | Global iteration |
| `This _.coll actor` | Collection items | New (fresh) | Collection iteration |

**Could Work Better**: Make delegation explicit:

```
Du actor                    # Delegate, same context
Du.new actor               # Delegate, fresh frame
Invoke actor with ${data}  # Delegate with explicit data
Call actor arg1 arg2       # Delegate with positional args
```

---

## Synthesis: The Generalized Code Generation Machine

Based on these mechanisms, here's a unified model:

```
┌─────────────────────────────────────────────────────────────┐
│                    CODE GENERATION ENGINE                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────┐     ┌─────────────┐     ┌─────────────┐   │
│  │   SCHEMA    │────→│   ACTORS    │────→│   OUTPUT    │   │
│  │  (Data)     │     │  (Rules)    │     │  (Code)     │   │
│  └─────────────┘     └─────────────┘     └─────────────┘   │
│         │                   │                   ↑           │
│         │                   │                   │           │
│         ▼                   ▼                   │           │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                  CONTEXT STACK                       │   │
│  │                                                      │   │
│  │  [0] Root        ─────────────────────────────────  │   │
│  │  [1] All Comp    ← foreach component                │   │
│  │  [2] gen_comp    ← pattern match: find=Find         │   │
│  │  [3] Its Element ← foreach element                  │   │
│  │  [4] gen_field   ← pattern match: type=ref          │   │
│  │       ↓                                              │   │
│  │       C "K${name}p int" ──────────────────→ OUTPUT  │   │
│  │                                                      │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                PATH RESOLUTION                       │   │
│  │                                                      │   │
│  │  ${name}           → current.Names["name"]          │   │
│  │  ${parent.name}    → current.parent.Names["name"]   │   │
│  │  ${.gen_comp.name} → stack[gen_comp].Names["name"]  │   │
│  │  ${_.map.key}      → globals["map"]["key"]          │   │
│  │  ${name:c}         → capitalize(resolve("name"))    │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │              CONDITIONAL EXECUTION                   │   │
│  │                                                      │   │
│  │  Actor name Comp attr op value                      │   │
│  │    └─ All actors with 'name' tried                  │   │
│  │    └─ Guards evaluated: attr op value               │   │
│  │    └─ All matching actors execute                   │   │
│  │                                                      │   │
│  │  Operators: = != in !in has regex ?? ? & |          │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  ┌─────────────────────────────────────────────────────┐   │
│  │               DEFERRED EXECUTION                     │   │
│  │                                                      │   │
│  │  Out delay    → Record but don't emit               │   │
│  │  trig()       → Replay when child emits             │   │
│  │  reGoCmds()   → Re-execute delayed commands         │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Key Architectural Insights

### 1. **Separation of Concerns**
- Schema defines **what exists** (components, relationships)
- Actors define **what to generate** (patterns, conditions)
- Runtime manages **how to execute** (stack, resolution)

### 2. **Declarative Over Imperative**
The system avoids explicit control flow in favor of:
- Pattern matching (actor conditions)
- Implicit iteration (All, Its, This)
- Lazy evaluation (Out delay, trig)

### 3. **Context is Everything**
Every operation happens within nested contexts. The Win stack is not just a call stack—it's **addressable working memory** that any expression can query.

### 4. **Multi-Pass as Natural Pattern**
Reference resolution requires multiple passes because dependencies form DAGs, not trees. The system handles this by:
- Tracking error counts
- Iterating until convergence
- Allowing partial resolution

### 5. **Composition Over Inheritance**
Actors compose via:
- `Du` (delegation)
- `Its`/`All`/`This` (iteration)
- Chained conditions (`&`, `|`)
- Named context access (`${.name.field}`)


# Break Command: Deep Analysis

## Current Implementation

```go
case *KpBreak:
    // Conditional execution based on IsCheck flag
    if c.Kcheck == "True" && !glob.Wins[winp].IsCheck {
        continue
    }
    if c.Kcheck == "False" && glob.Wins[winp].IsCheck {
        continue
    }
    
    ret := 0
    if c.Kwhat == "E_O_L" || c.Kwhat == "actor" {
        ret = 2
    }
    if c.Kwhat == "loop" {
        ret = 1
    }
    if c.Kwhat == "cmds" {
        ret = 3
    }
    if c.Kwhat == "exit" {
        if glob.LoadErrs > 0 || glob.RunErrs > 0 {
            return -4
        }
        // BUG: Falls through with ret=0 if no errors!
    }
    
    // Named actor break - negates return and sets flag
    if c.Kactor != "E_O_L" && c.Kactor != "." {
        for i := winp - 1; i >= 0; i-- {
            if glob.Wins[i].Name == c.Kactor {
                glob.Wins[i+1].BrkAct = true
                ret = -ret
                break
            }
        }
    }
    return ret
```

---

## Return Code Flow Through the System

### In `goCmds` → Back to `GoAct`

```go
// goCmds returns ret to GoAct
ret := goCmds(glob, i, winp)

if ret == -4 {
    return ret           // Exit: propagate immediately
}
if ret == 0 || ret == 3 {
    continue             // Success or cmds: try next actor
}

nret := ret
if ret == 2 {
    nret = 0             // Actor break: stop matching, return success
}
if ret < 0 && glob.Wins[glob.Winp].BrkAct {
    nret = -ret          // Named break: convert back to positive
}
glob.Winp--
return nret
```

### Complete Return Code Matrix

| `Kwhat` | Initial `ret` | After named actor | In GoAct becomes | Effect |
|---------|---------------|-------------------|------------------|--------|
| `actor` / `E_O_L` | 2 | -2 if named | 0 (stop) | Stop actor matching, continue loop |
| `loop` | 1 | -1 if named | 1 (propagate) | Stop actor iteration and loop |
| `cmds` | 3 | -3 if named | continue | Stop cmds, continues to next actor |
| `exit` (with errors) | -4 | n/a | -4 (propagate) | Terminate everything |
| `exit` (no errors) | 0 | 0 if named | continue | **BUG: does nothing** |

---

## The Named Actor Break Mechanism

```go
if c.Kactor != "E_O_L" && c.Kactor != "." {
    for i := winp - 1; i >= 0; i-- {
        if glob.Wins[i].Name == c.Kactor {
            glob.Wins[i+1].BrkAct = true  // Mark the frame AFTER target
            ret = -ret                     // Negate return code
            break
        }
    }
}
```

**Key Insight**: Named breaks work by:
1. Walking **backward** through Win stack
2. Finding the named actor's frame
3. Setting `BrkAct` on the **next** frame (i+1)
4. Negating `ret` so it propagates up

### Why `i+1` Not `i`?

```
Win Stack:
[3] gen_field  ← Current (winp=3), Break actor outer_loop
[2] gen_elem   ← BrkAct set HERE (i+1)
[1] outer_loop ← Target found HERE (i)
[0] main
```

The flag is set on `i+1` because that's the frame that will **check** the flag when unwinding. The target frame (`outer_loop`) doesn't need the flag—it's the frames **between** current and target that need to know to propagate.

### Propagation Logic

```go
// In GoAct, when ret < 0:
if ret < 0 && glob.Wins[glob.Winp].BrkAct {
    nret = -ret  // Convert -2 back to 2, -1 back to 1
}
```

Each frame checks: "Am I marked for break propagation?"
- If yes: flip sign back and return
- This continues until reaching the target

---

## The `IsCheck` Gate

```go
if c.Kcheck == "True" && !glob.Wins[winp].IsCheck {
    continue  // Skip break if IsCheck is false
}
if c.Kcheck == "False" && glob.Wins[winp].IsCheck {
    continue  // Skip break if IsCheck is true
}
```

**Purpose**: Conditional break based on some validation state.

**Problem**: `IsCheck` is never set in the visible code! It's likely set by:
- `Add.check` operations (checks for duplicates)
- Validation commands
- But the setting mechanism is in disabled/missing code

---

## Identified Bugs and Edge Cases

### Bug 1: `Break exit` Without Errors Does Nothing

```go
if c.Kwhat == "exit" {
    if glob.LoadErrs > 0 || glob.RunErrs > 0 {
        return -4
    }
    // Falls through! ret still 0
}
```

**Expected**: `Break exit` should always exit
**Actual**: Only exits if errors exist

**Fix**:
```go
if c.Kwhat == "exit" {
    if glob.LoadErrs > 0 || glob.RunErrs > 0 {
        return -4
    }
    return -5  // Clean exit code
}
```

### Bug 2: `Break cmds` Continues to Next Actor

```go
// In GoAct:
if ret == 0 || ret == 3 {
    continue  // ret=3 (cmds) treated same as success!
}
```

**Expected**: `Break cmds` should stop current actor's commands only
**Actual**: Stops commands AND continues to next actor (same as success)

**The Semantic Confusion**:
- `Break actor` = "I'm done, don't try other actors"
- `Break cmds` = "I'm done with THIS actor's commands" → should it try other actors?

The current behavior says **yes**, as intended

```

### Bug 3: Named Break on Non-Existent Actor

```go
for i := winp - 1; i >= 0; i-- {
    if glob.Wins[i].Name == c.Kactor {
        // Found
        break
    }
}
// If not found, ret stays positive, BrkAct never set
```

**Expected**: Error or warning if target actor not in stack
**Actual**: Silent failure, break has no effect

---

## Complete Break Semantics

### Current (With Bugs)

```
┌─────────────────────────────────────────────────────────────┐
│                      BREAK COMMAND                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Break [check] [what] [actor]                               │
│                                                             │
│  check:                                                     │
│    (none)  → Always execute                                 │
│    True    → Only if IsCheck is true                        │
│    False   → Only if IsCheck is false                       │
│                                                             │
│  what:                                                      │
│    actor   → ret=2  → Stop matching actors for this name    │
│    loop    → ret=1  → Exit current Its/All/This iteration   │
│    cmds    → ret=3  → Stop commands continues actors │
│    exit    → ret=-4 → Terminate (BUG: only if errors)       │
│    (none)  → ret=2  → Same as "actor"                       │
│                                                             │
│  actor:                                                     │
│    (none)  → Break applies to immediate context             │
│    .       → Same as none                                   │
│    name    → Break propagates up to named actor             │
│             └─ Negates ret, sets BrkAct flags               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Corrected Semantics

```
┌─────────────────────────────────────────────────────────────┐
│                   CORRECTED BREAK COMMAND                   │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Break [what] [target]                                      │
│                                                             │
│  what (scope of break):                                     │
│    actor   → Stop trying actors, return to caller           │
│    loop    → Exit current iteration (Its/All/This)          │
│    cmds    → Stop commands, continue actors            │
│    exit    → Terminate entire generation                    │
│    error   → Terminate with error status                    │
│                                                             │
│  target (where to break TO):                                │
│    (none)  → Immediate enclosing scope                      │
│    .       → Same as none                                   │
│    name    → Named actor in Win stack                       │
│                                                             │
│  Return codes:                                              │
│    0  = Success, continue                                   │
│    1  = Loop break                                          │
│    2  = Actor break                                         │
│    3  = Cmds break, continue to next actor      │
│   -4  = Error exit                                          │
│   -5  = Clean exit                                          │
│   -N  = Propagating break (will be flipped at target)       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Generalized Break Architecture

### The Control Flow Model

```
┌─────────────────────────────────────────────────────────────┐
│                    EXECUTION LAYERS                         │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Layer 4: Generation          Break exit                    │
│           ↑                      ↓                          │
│  Layer 3: Actor Matching      Break actor                   │
│           ↑                      ↓                          │
│  Layer 2: Command Execution   Break cmds                    │
│           ↑                      ↓                          │
│  Layer 1: Iteration           Break loop                    │
│           ↑                      ↓                          │
│  Layer 0: Individual Step     (no break, just return)       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

Each break type exits to a specific layer:

| Break Type | Exits From | Returns To | Continues? |
|------------|------------|------------|------------|
| `loop` | Current iteration | Iteration control | Next iteration |
| `cmds` | Command list | Command iteration | Next  |
| `actor` | Actor matching | Caller (Its/All/Du) | Normal return, end loop |
| `exit` | Entire stack | OS | No |

### Named Break as Exception Handling

Named breaks are like **labeled breaks** in Go or **exceptions** caught at specific frames:

```go
// Conceptual equivalent:
outer_loop:
for _, comp := range components {
    for _, elem := range comp.Elements {
        if elem.Type == "error" {
            break outer_loop  // Jump to labeled loop
        }
    }
}
```

The Win stack implementation:
```
[5] Check error    → Break loop outer_loop
[4] gen_elem       → BrkAct=true, propagate -1
[3] inner_loop     → BrkAct=true, propagate -1
[2] gen_comp       → BrkAct=true, propagate -1
[1] outer_loop     → Name matches! Flip to +1, handle loop break
[0] main
```

---

## Proposed Generalized Break System

### Clear Semantics

```
Break scope [target] [condition]

scope:
  step     → Return from current command only ?
  cmds     → Return from current actor's command list
  actor    → Return from actor matching loop
  frame    → Return from current Win frame ?
  loop     → Exit iteration, continue to next
  all      → Exit iteration completely ?
  exit     → Terminate generation (clean)
  error    → Terminate generation (with error)

target:
  .        → Immediate scope
  name     → Propagate up to named actor
  @N       → Propagate up N frames
  @root    → Propagate to root

condition:
  (none)   → Always break
  if expr  → Break only if expression true
  unless e → Break only if expression false
  check    → Break based on IsCheck flag
```

### Example Patterns

```
# Simple breaks
Break loop                    # Next iteration
Break actor                   # Done with this actor
Break exit                    # Clean termination

# Named breaks
Break loop outer              # Exit to named actor's loop
Break actor gen_comp          # Stop gen_comp's actor matching

# Conditional breaks
Break actor if ${count} > 10  # Break if condition met
Break loop unless ${valid}    # Break unless valid

# Scoped breaks
Break frame @2                # Exit 2 frames up
Break all @root               # Unwind entire stack
```

### Return Code Redesign

```go
type BreakCode int

const (
    Continue    BreakCode = 0   // Normal continuation
    BreakLoop   BreakCode = 1   // Exit current iteration
    BreakActor  BreakCode = 2   // Exit actor matching
    BreakCmds   BreakCode = 3   // Exit command list
    BreakFrame  BreakCode = 4   // Exit Win frame
    ExitClean   BreakCode = -1  // Clean termination
    ExitError   BreakCode = -2  // Error termination
    
    // Propagating (negative of base + 100)
    PropLoop    BreakCode = -101
    PropActor   BreakCode = -102
    PropCmds    BreakCode = -103
    PropFrame   BreakCode = -104
)

func (b BreakCode) IsPropagating() bool {
    return b <= -100
}

func (b BreakCode) BaseCode() BreakCode {
    if b <= -100 {
        return -(b + 100)
    }
    return b
}
```

---

## The Missing Piece: `IsCheck`

Looking at the code, `IsCheck` is checked but never set in `goCmds`:

```go
if c.Kcheck == "True" && !glob.Wins[winp].IsCheck {
    continue
}
```

This suggests there should be commands like:

```
Add.break _.seen:${name} ${value}
# Same as Break actor if duplicate found, no need for break cmd - for most cases

Add.check _.seen:${name} ${value}
# Sets IsCheck = true if duplicate found

Break loop . True
# Only breaks if IsCheck is true (duplicate detected)
# This is for when break variations (loop,cmds) are needed.

Add.check _.seen:${name} ${value}
Break actor . False
# Only breaks if IsCheck is false (no duplicate - new)
# This is for new entries.

Add.check.no_add _.seen:${name} ${value}
# Check only

Add.check.or _.seen:${name} ${value}
Add.check.and _.seen:${name} ${value}
# Future options
```

### Reconstructed Check Pattern

```go
case *KpAdd:
    if slices.Contains(c.Flags, "check") {
        // Check if key already exists
        if _, exists := collection[key]; exists {
            glob.Wins[winp].IsCheck = true  // Mark duplicate
        }
    }
    // ... rest of add logic
```

**Usage**:
```
Actor collect_unique Element
  Add.check _.seen:${name} ${value}
  Break check actor    # Skip if duplicate
  C Processing ${name}
```

---

## Visual: Break Flow Through Stack

```
Initial state:                After Break loop outer:
                             
[4] validate    ─┐           [4] validate    ─┐
    Break loop   │               ret = -1     │ Propagating
    outer        │               BrkAct=true  │
                 │                            │
[3] gen_elem    ─┤           [3] gen_elem    ─┤
    Its...       │               ret = -1     │ Propagating
                 │               BrkAct=true  │
                 │                            │
[2] inner       ─┤           [2] inner       ─┤
    All...       │               ret = -1     │ Propagating
                 │               BrkAct=true  │
                 │                            │
[1] outer       ─┘           [1] outer       ─┘
    All Comp     ← Target        ret = +1     ← Caught!
                                 Handle loop   
                                 break        
                             
[0] main                     [0] main         
                                 Continues    
```

The negative return propagates up, each frame's `BrkAct` flag confirms it should continue propagating, until the target frame sees its name matches and flips the sign back to handle locally.

