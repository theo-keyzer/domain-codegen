# From Batch Generator to Continuous Multi-Threaded Database Engine

## Current Architecture vs. Target

```
CURRENT                              TARGET
────────────────────────────────────────────────────────────────
Single GlobT                    →    Pool of Session contexts
Single Win stack                →    Per-session stacks + shared state
Linear execution                →    Concurrent actor evaluation
Batch output (stdout)           →    Multiple output streams/subscribers
Run once, exit                  →    Continuous daemon
Load from files                 →    Live database backing
Manual restart for changes      →    Hot reload / self-update
```

**Answer: Yes, it's possible.** The existing primitives are surprisingly close to what's needed. Here's the evolution path:

---

## Phase 1: Decouple State Layers

### Current Problem
```go
type GlobT struct {
    Acts     *ActT           // Actors (code)
    Dats     *ActT           // Data (schema instances)
    Wins     []*WinT         // Execution stack (single thread!)
    Collect  map[string]any  // Runtime collections
    OutOn    bool            // Global output state
    // ... all mixed together
}
```

### Solution: Layered State Model

```go
// Shared across all sessions (read-mostly)
type SchemaLayer struct {
    sync.RWMutex
    Acts     *ActT                    // Actor definitions
    Dats     *ActT                    // Schema/data definitions
    Version  uint64                   // For change detection
    Index    map[string]int           // Global lookups
}

// Per-session execution context
type SessionT struct {
    ID       string
    Schema   *SchemaLayer             // Shared reference
    Wins     []WinT                   // Private stack
    Collect  map[string]any           // Session-local collections
    Output   io.Writer                // Session-specific output
    Subs     []chan<- Event           // Subscribers to this session's events
}

// Cross-session persistent state
type DatabaseLayer struct {
    sync.RWMutex
    Tables   map[string]*Table        // Persistent collections
    WAL      *WriteAheadLog           // Durability
    Triggers map[string][]TriggerDef  // Data change triggers
}

// The new "global" coordinator
type EngineT struct {
    Schema   *SchemaLayer
    Database *DatabaseLayer
    Sessions sync.Map                 // map[string]*SessionT
    Events   chan Event               // Central event bus
    Done     chan struct{}
}
```

---

## Phase 2: Multi-Session Execution

### Session Isolation

Each query/request gets its own session:

```go
func (engine *EngineT) NewSession(id string, output io.Writer) *SessionT {
    return &SessionT{
        ID:      id,
        Schema:  engine.Schema,           // Shared read
        Wins:    make([]WinT, 0, 64),     // Private stack
        Collect: make(map[string]any),    // Private collections
        Output:  output,
    }
}

func (engine *EngineT) Execute(sess *SessionT, actorName string, data any) error {
    NewAct(sess, actorName, "", "api")
    return GoAct(sess, data)
}
```

### Modified GoAct for Sessions

```go
// Change: glob *GlobT → sess *SessionT
func GoAct(sess *SessionT, dat interface{}) int {
    winp := len(sess.Wins)
    sess.Wins = append(sess.Wins, WinT{Dat: dat, Name: sess.Wins[winp-1].Name})
    
    // Read lock for actor iteration (allows concurrent reads)
    sess.Schema.RLock()
    actors := sess.Schema.Acts.ApActor
    sess.Schema.RUnlock()
    
    for i := 0; i < len(actors); i++ {
        // ... same matching logic, but on sess not glob
    }
}
```

---

## Phase 3: Persistent Collections (The Database)

### Extend `Add` Command for Persistence

```go
type AddTarget int

const (
    AddSession   AddTarget = iota  // _. prefix - session local
    AddDatabase                     // @. prefix - persistent
    AddBroadcast                    // !. prefix - emit event
)

func parseAddPath(path string) (AddTarget, string) {
    switch {
    case strings.HasPrefix(path, "_."):
        return AddSession, path[2:]
    case strings.HasPrefix(path, "@."):
        return AddDatabase, path[2:]
    case strings.HasPrefix(path, "!."):
        return AddBroadcast, path[2:]
    default:
        return AddSession, path
    }
}
```

### Database-Backed Collections

```go
func addCmd(sess *SessionT, winp int, c *KpAdd, lno string) int {
    target, path := parseAddPath(c.Kpath)
    
    switch target {
    case AddSession:
        // Current behavior - session local
        sess.Collect[path] = value
        
    case AddDatabase:
        // Persistent write with transaction
        sess.Engine.Database.Lock()
        defer sess.Engine.Database.Unlock()
        
        sess.Engine.Database.Tables[table].Put(key, value)
        sess.Engine.Database.WAL.Append(PutOp{table, key, value})
        
        // Trigger any watchers
        sess.Engine.fireTriggers(table, key, value)
        
    case AddBroadcast:
        // Emit event to all subscribers
        sess.Engine.Events <- Event{
            Type:    "add",
            Path:    path,
            Value:   value,
            Session: sess.ID,
        }
    }
    return 0
}
```

---

## Phase 4: Continuous Operation (Event Loop)

### The Engine Daemon

```go
func (engine *EngineT) Run(ctx context.Context) error {
    // Start background workers
    for i := 0; i < runtime.NumCPU(); i++ {
        go engine.worker(ctx, i)
    }
    
    // Event loop
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
            
        case event := <-engine.Events:
            engine.handleEvent(event)
            
        case change := <-engine.Schema.Changes:
            engine.hotReload(change)
        }
    }
}

func (engine *EngineT) handleEvent(event Event) {
    switch event.Type {
    case "query":
        // Spawn session for query
        sess := engine.NewSession(event.ID, event.Output)
        go func() {
            engine.Execute(sess, event.Actor, event.Data)
            event.Done <- struct{}{}
        }()
        
    case "schema_change":
        // Hot reload actors/data
        engine.hotReload(event)
        
    case "data_change":
        // Re-trigger affected actors
        engine.retrigger(event)
    }
}
```

### Query Interface

```go
// External API
func (engine *EngineT) Query(actor string, data any) (string, error) {
    var buf strings.Builder
    done := make(chan struct{})
    
    engine.Events <- Event{
        Type:   "query",
        Actor:  actor,
        Data:   data,
        Output: &buf,
        Done:   done,
    }
    
    <-done
    return buf.String(), nil
}

// Example usage:
result, _ := engine.Query("gen_component", componentData)
```

---

## Phase 5: Self-Updating (Hot Reload)

### Watch for Changes

```go
func (engine *EngineT) watchFiles(ctx context.Context, paths []string) {
    watcher, _ := fsnotify.NewWatcher()
    for _, p := range paths {
        watcher.Add(p)
    }
    
    for {
        select {
        case <-ctx.Done():
            return
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write != 0 {
                engine.Events <- Event{
                    Type: "schema_change",
                    Path: event.Name,
                }
            }
        }
    }
}

func (engine *EngineT) hotReload(change Event) {
    // Parse changed file
    newActs, newDats := load(change.Path)
    
    // Atomic swap with version bump
    engine.Schema.Lock()
    engine.Schema.Acts = newActs
    engine.Schema.Dats = newDats
    engine.Schema.Version++
    engine.Schema.Unlock()
    
    // Notify all sessions
    engine.Sessions.Range(func(key, value any) bool {
        sess := value.(*SessionT)
        sess.Notify(SchemaChanged{Version: engine.Schema.Version})
        return true
    })
}
```

### Self-Modifying Actors

Re-enable `New` command for runtime schema modification:

```go
case *KpNew:
    // Create new component at runtime
    _, line := strs(sess, winp, c.KLine, c.LineNo, true, true)
    
    if c.KWhere == "acts" {
        // Adding new actor - requires write lock
        sess.Engine.Schema.Lock()
        load(sess.Engine.Schema.Acts, c.KWhat, line, 0, c.LineNo)
        sess.Engine.Schema.Version++
        sess.Engine.Schema.Unlock()
        
    } else if c.KWhere == "dats" {
        // Adding new data - can be session or persistent
        if c.Persistent {
            sess.Engine.Database.Lock()
            load(sess.Engine.Database.Tables[c.KWhat], line)
            sess.Engine.Database.Unlock()
        } else {
            load(sess.LocalDats, c.KWhat, line, 0, c.LineNo)
        }
    }
```

---

## Phase 6: Reactive Triggers

### Data Change Triggers

```go
type TriggerDef struct {
    Table    string   // Watch this table
    Actor    string   // Run this actor
    On       string   // "insert" | "update" | "delete" | "*"
    Filter   string   // Optional condition
}

func (engine *EngineT) fireTriggers(table, key string, value any) {
    triggers := engine.Database.Triggers[table]
    
    for _, trig := range triggers {
        // Spawn session for each trigger
        sess := engine.NewSession("trigger-"+uuid(), io.Discard)
        sess.Collect["_trigger"] = map[string]any{
            "table": table,
            "key":   key,
            "value": value,
            "op":    "insert",
        }
        
        go engine.Execute(sess, trig.Actor, value)
    }
}
```

### Actor-Defined Triggers

```
# Define trigger in actor file
Trigger on_user_insert @.users insert
  Actor on_user_insert User
    Du validate_user
    Du index_user
    Add @.user_index:${email} ${id}
    Add !.events { "type": "user_created", "id": "${id}" }
```

---

## Phase 7: The Win Stack in Concurrent Context

### Problem: Triggers/Wormholes Don't Work Across Sessions

The `trig()` function walks backward through a single Win stack. In a multi-session world, we need:

1. **Session-local wormholes**: Current behavior, within one session
2. **Cross-session signals**: Events, not wormholes

```go
func trig(sess *SessionT, winp int) {
    // Session-local trigger (unchanged logic)
    if !sess.Wins[winp].IsPrev || winp == 0 {
        return
    }
    // ... same as before, but on sess.Wins
}

// NEW: Cross-session notification
func notify(sess *SessionT, target string, data any) {
    sess.Engine.Events <- Event{
        Type:    "notify",
        Target:  target,     // Session ID or "*" for broadcast
        Data:    data,
        Source:  sess.ID,
    }
}
```

### Concurrent Actor Execution Within Session

For parallelizing `All Comp gen_comp`:

```go
case *KpAll:
    // Get all components of type
    comps := getAllOfType(sess, c.Kwhat)
    
    if c.Parallel && len(comps) > 1 {
        // Parallel execution
        var wg sync.WaitGroup
        results := make(chan int, len(comps))
        
        for _, comp := range comps {
            wg.Add(1)
            go func(c any) {
                defer wg.Done()
                // Each gets a FORKED session
                fork := sess.Fork()
                NewAct(fork, c.Kactor, args, c.LineNo)
                results <- GoAct(fork, c)
            }(comp)
        }
        
        wg.Wait()
        close(results)
        
        // Merge results
        for ret := range results {
            if ret != 0 {
                return ret
            }
        }
    } else {
        // Sequential (current behavior)
        for _, comp := range comps {
            NewAct(sess, c.Kactor, args, c.LineNo)
            ret := GoAct(sess, comp)
            if ret != 0 {
                return ret
            }
        }
    }
```

---

## Complete Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                         ENGINE (Daemon)                             │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐          │
│  │   SCHEMA     │    │   DATABASE   │    │    EVENT     │          │
│  │   LAYER      │    │    LAYER     │    │     BUS      │          │
│  │              │    │              │    │              │          │
│  │  Acts        │    │  Tables      │    │  Queries     │          │
│  │  Dats        │    │  WAL         │    │  Changes     │          │
│  │  Index       │    │  Triggers    │    │  Triggers    │          │
│  │  Version     │    │  Snapshots   │    │  Broadcasts  │          │
│  └──────┬───────┘    └──────┬───────┘    └──────┬───────┘          │
│         │ read              │ read/write        │ pub/sub          │
│         └─────────┬─────────┴─────────┬─────────┘                  │
│                   │                   │                            │
│         ┌─────────▼─────────┐ ┌───────▼─────────┐                  │
│         │    SESSION 1      │ │    SESSION 2    │ ...              │
│         │                   │ │                 │                  │
│         │  Wins[] (stack)   │ │  Wins[] (stack) │                  │
│         │  Collect (local)  │ │  Collect (local)│                  │
│         │  Output (writer)  │ │  Output (writer)│                  │
│         └─────────┬─────────┘ └───────┬─────────┘                  │
│                   │                   │                            │
│         ┌─────────▼───────────────────▼─────────┐                  │
│         │            WORKER POOL                │                  │
│         │                                       │                  │
│         │  GoAct() executions on sessions       │                  │
│         │  Parallel All/Its when flagged        │                  │
│         │  Trigger evaluations                  │                  │
│         └───────────────────────────────────────┘                  │
│                                                                     │
├─────────────────────────────────────────────────────────────────────┤
│                         EXTERNAL INTERFACES                         │
│                                                                     │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐            │
│  │   HTTP   │  │   gRPC   │  │  WebSock │  │   CLI    │            │
│  │   API    │  │   API    │  │  Stream  │  │  REPL    │            │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └────┬─────┘            │
│       └─────────────┴─────────────┴─────────────┘                  │
│                            │                                        │
│                     Events Channel                                  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## New Commands Needed

### Session/Database Commands

```
# Session-local (current behavior, explicit)
Add _.map:key value

# Persistent (database-backed)
Add @.table:key value
Get @.table:key target_var
Del @.table:key

# Broadcast event
Add !.channel { data }
Sub !.channel actor_name

# Concurrency control
Lock @.table timeout_ms
Unlock @.table
Transaction
  Add @.a:1 x
  Add @.b:2 y
Commit
```

### Control Flow Extensions

```
# Parallel iteration
All.parallel Comp gen_comp

# Async execution
Async actor_name args
Await result_var

# Watch for changes
Watch @.users on_user_change
Unwatch @.users

# Session management
Fork child_session_id
Join child_session_id result_var
```

---

## Migration Path

### Step 1: Extract Session (Low Risk)
```go
// Replace glob with sess in all functions
// Sessions pool with shared Schema
// Single-threaded initially
```

### Step 2: Add Database Layer (Medium Risk)
```go
// Keep collections in-memory initially
// Add WAL for persistence
// @. prefix for persistent adds
```

### Step 3: Event Loop (Medium Risk)
```go
// Engine.Run() daemon
// Query interface
// File watchers for hot reload
```

### Step 4: Multi-Threading (Higher Risk)
```go
// Worker pool for sessions
// Parallel All/Its
// Lock management for writes
```

### Step 5: Reactive Triggers (Medium Risk)
```go
// Watch/trigger infrastructure
// Cross-session notifications
// Event subscriptions
```

---

## Example: Continuous Schema Evolution

```
# Self-evolving system that watches its own schema
Actor main .
  Watch @.components on_component_change
  Watch @.actors on_actor_change
  All @.components gen_code

Actor on_component_change Component
  Du validate_component
  Du regenerate_affected
  Add !.updates { "type": "component", "name": "${name}" }

Actor on_actor_change Actor
  Du validate_actor
  Du hot_reload_actor
  Add !.updates { "type": "actor", "name": "${name}" }

Actor regenerate_affected Component
  # Find all actors that reference this component
  This @.actor_deps:${name} trigger_regen

Actor trigger_regen Actor
  # Re-execute this actor's generation
  Du ${name}
```

---

## The Self-Updating Eternal Loop

From the canons, realized:

```
┌────────────────────────────────────────────────────────────┐
│                    ETERNAL RUNNER v2                       │
├────────────────────────────────────────────────────────────┤
│                                                            │
│  Engine.Run():                                             │
│    for {                                                   │
│      1. Check for schema changes                           │
│         → Hot reload if changed                            │
│                                                            │
│      2. Check for data changes                             │
│         → Fire triggers                                    │
│         → Re-execute affected actors                       │
│                                                            │
│      3. Process query queue                                │
│         → Spawn sessions                                   │
│         → Execute actors                                   │
│         → Return results                                   │
│                                                            │
│      4. Evaluate golden promotions                         │
│         → If new golden found                              │
│         → Replace self (the runner)                        │
│         → Continue with new code                           │
│                                                            │
│      5. Persist state                                      │
│         → Checkpoint WAL                                   │
│         → Snapshot if needed                               │
│    }                                                       │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

**This is achievable.** The current codebase is ~70% of the way there. The key missing pieces are:

1. **Session isolation** (2-3 days of refactoring)
2. **Persistence layer** (1 week)
3. **Event loop** (2-3 days)
4. **Concurrency primitives** (1 week)
5. **Hot reload** (2-3 days)

The actor model is already well-suited to this evolution. The declarative pattern matching can be parallelized. The Win stack is naturally session-scoped. The path resolution is already a query language.
