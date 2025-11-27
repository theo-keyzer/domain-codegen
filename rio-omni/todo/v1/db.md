# Append-Only Database with Tail Updates and Batch Linking

This is exactly what the Rio loader already does. The architecture just needs to be made persistent and continuous.

---

## Current Pattern (Already Working)

```
┌─────────────────────────────────────────────────────────────┐
│                  RIO LOADER = YOUR DATABASE                 │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Phase 1: APPEND                                            │
│  ─────────────────                                          │
│  load(act, "Comp", line, ...)                               │
│    → act.ApComp = append(act.ApComp, st)                    │
│    → act.index["Comp_"+name] = st.Me                        │
│                                                             │
│  Phase 2: LINK                                              │
│  ─────────────────                                          │
│  refs(act)                                                  │
│    → for each component with unresolved refs                │
│    →   lookup in index                                      │
│    →   set Kfieldp pointer                                  │
│    → repeat until converged                                 │
│                                                             │
│  Phase 3: OUTPUT                                            │
│  ─────────────────                                          │
│  GoAct(glob, data)                                          │
│    → traverse linked structures                             │
│    → emit generated code                                    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## Extended for Continuous Batch Operation

### Core Data Structure

```go
type DatabaseT struct {
    sync.RWMutex
    
    // The append-only arrays (your existing structure)
    Tables map[string]*TableT
    
    // Batch tracking
    CurrentBatch  uint64
    BatchBoundary map[string]int  // table -> index where current batch starts
    
    // Persistence
    WAL      *os.File
    Segments []Segment  // Immutable historical segments
}

type TableT struct {
    Name   string
    Items  []Kp              // Append-only array
    Index  map[string]int    // Key -> position
    Schema *KpComp           // Component definition
    
    // Tail tracking
    TailStart    int         // Where "mutable" region begins
    UnresolvedRefs []int     // Items needing resolution
}

type Segment struct {
    ID        uint64
    StartIdx  int
    EndIdx    int
    Frozen    bool          // No more updates allowed
    FilePath  string        // Persistent storage
}
```

---

## Append-Only with Tail Updates

### The Invariant

```
┌─────────────────────────────────────────────────────────────┐
│                      TABLE STRUCTURE                        │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Index:    0    100    200    300    400    500             │
│            │     │      │      │      │      │              │
│            ▼     ▼      ▼      ▼      ▼      ▼              │
│  Items: [=====|=====|=====|=====|~~~~~|~~~~~]               │
│          ─────────────────────── ───────────                │
│               FROZEN              TAIL                      │
│            (immutable)         (mutable)                    │
│                                                             │
│  Segment 1  Seg 2  Seg 3  Seg 4   Current Batch             │
│  (on disk) (disk) (disk) (disk)   (in memory)               │
│                                                             │
│  Rules:                                                     │
│  • Append: Always to end of Items                           │
│  • Update: Only items where idx >= TailStart                │
│  • Refs: Can point anywhere (even to frozen segments)       │
│  • Freeze: TailStart moves forward, segment written         │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Implementation

```go
// Append new item (always allowed)
func (t *TableT) Append(item Kp) int {
    idx := len(t.Items)
    t.Items = append(t.Items, item)
    
    // Index by key if applicable
    if key := item.GetKey(); key != "" {
        t.Index[key] = idx
    }
    
    // Track for resolution
    if item.HasUnresolvedRefs() {
        t.UnresolvedRefs = append(t.UnresolvedRefs, idx)
    }
    
    return idx
}

// Update item (only in tail)
func (t *TableT) Update(idx int, field string, value any) error {
    if idx < t.TailStart {
        return fmt.Errorf("cannot update frozen item at %d (tail starts %d)", 
                         idx, t.TailStart)
    }
    
    t.Items[idx].SetField(field, value)
    return nil
}

// Update reference pointer (only in tail)
func (t *TableT) SetRef(idx int, field string, targetIdx int) error {
    if idx < t.TailStart {
        return fmt.Errorf("cannot update ref in frozen item %d", idx)
    }
    
    // Target can be anywhere (even frozen)
    t.Items[idx].SetRefPointer(field, targetIdx)
    return nil
}

// Freeze current tail into segment
func (t *TableT) Freeze() *Segment {
    seg := &Segment{
        ID:       nextSegmentID(),
        StartIdx: t.TailStart,
        EndIdx:   len(t.Items),
        Frozen:   true,
    }
    
    // Move tail forward
    t.TailStart = len(t.Items)
    t.UnresolvedRefs = nil  // Clear - should be empty after resolution
    
    return seg
}
```

---

## Batch Append with Linking

### The Batch Cycle

```
┌─────────────────────────────────────────────────────────────┐
│                     BATCH LIFECYCLE                         │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌─────────────┐                                            │
│  │  RECEIVE    │  ← Batch of raw records arrives            │
│  │  BATCH      │                                            │
│  └──────┬──────┘                                            │
│         │                                                   │
│         ▼                                                   │
│  ┌─────────────┐                                            │
│  │  APPEND     │  ← Add to arrays, build local indices      │
│  │  PHASE      │  ← All refs initialized to -1              │
│  └──────┬──────┘                                            │
│         │                                                   │
│         ▼                                                   │
│  ┌─────────────┐                                            │
│  │  LINK       │  ← Multi-pass reference resolution         │
│  │  PHASE      │  ← Can reference frozen OR current batch   │
│  └──────┬──────┘                                            │
│         │                                                   │
│         ▼                                                   │
│  ┌─────────────┐                                            │
│  │  VALIDATE   │  ← Check all refs resolved                 │
│  │  PHASE      │  ← Check constraints                       │
│  └──────┬──────┘                                            │
│         │                                                   │
│    ┌────┴────┐                                              │
│    │         │                                              │
│    ▼         ▼                                              │
│  ┌─────┐  ┌──────┐                                          │
│  │ OK  │  │ FAIL │                                          │
│  └──┬──┘  └──┬───┘                                          │
│     │        │                                              │
│     ▼        ▼                                              │
│  ┌──────┐ ┌────────┐                                        │
│  │FREEZE│ │ROLLBACK│  ← Remove batch items                  │
│  │BATCH │ │ BATCH  │                                        │
│  └──────┘ └────────┘                                        │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Implementation

```go
type BatchT struct {
    ID         uint64
    StartIndex map[string]int  // table -> starting index
    Items      map[string][]Kp // table -> items in this batch
    Status     BatchStatus
}

type BatchStatus int
const (
    BatchPending BatchStatus = iota
    BatchAppended
    BatchLinked
    BatchValidated
    BatchFrozen
    BatchRolledBack
)

func (db *DatabaseT) StartBatch() *BatchT {
    db.Lock()
    defer db.Unlock()
    
    batch := &BatchT{
        ID:         db.CurrentBatch + 1,
        StartIndex: make(map[string]int),
        Items:      make(map[string][]Kp),
        Status:     BatchPending,
    }
    
    // Record starting positions for potential rollback
    for name, table := range db.Tables {
        batch.StartIndex[name] = len(table.Items)
    }
    
    db.CurrentBatch = batch.ID
    return batch
}

func (db *DatabaseT) AppendBatch(batch *BatchT, table string, items []Kp) {
    db.Lock()
    defer db.Unlock()
    
    t := db.Tables[table]
    for _, item := range items {
        t.Append(item)
        batch.Items[table] = append(batch.Items[table], item)
    }
    
    batch.Status = BatchAppended
}

func (db *DatabaseT) LinkBatch(batch *BatchT) (int, error) {
    db.Lock()
    defer db.Unlock()
    
    totalErrs := 0
    prevErrs := 999999
    
    // Multi-pass resolution (your existing algorithm)
    for pass := 0; pass < 10; pass++ {
        errs := 0
        
        for tableName, table := range db.Tables {
            // Only resolve items in current batch's tail
            for _, idx := range table.UnresolvedRefs {
                if idx < batch.StartIndex[tableName] {
                    continue  // Not in this batch
                }
                
                item := table.Items[idx]
                errs += db.resolveRefs(item, idx)
            }
        }
        
        if errs == 0 {
            break  // All resolved
        }
        if errs == prevErrs {
            totalErrs = errs
            break  // Stuck
        }
        prevErrs = errs
    }
    
    if totalErrs == 0 {
        batch.Status = BatchLinked
    }
    return totalErrs, nil
}

func (db *DatabaseT) resolveRefs(item Kp, itemIdx int) int {
    errs := 0
    
    // For each reference field in the item
    for _, field := range item.RefFields() {
        if item.GetRefPointer(field) >= 0 {
            continue  // Already resolved
        }
        
        key := item.GetField(field).(string)
        targetTable := item.GetRefTable(field)
        
        // Lookup in index (can find frozen OR current items)
        if idx, ok := db.Tables[targetTable].Index[key]; ok {
            item.SetRefPointer(field, idx)
        } else {
            errs++
        }
    }
    
    return errs
}

func (db *DatabaseT) CommitBatch(batch *BatchT) error {
    if batch.Status != BatchLinked {
        return fmt.Errorf("batch not ready for commit: %v", batch.Status)
    }
    
    db.Lock()
    defer db.Unlock()
    
    // Write to WAL
    db.WAL.Write(serializeBatch(batch))
    db.WAL.Sync()
    
    // Freeze all tables at batch boundary
    for _, table := range db.Tables {
        seg := table.Freeze()
        db.persistSegment(seg)
    }
    
    batch.Status = BatchFrozen
    return nil
}

func (db *DatabaseT) RollbackBatch(batch *BatchT) {
    db.Lock()
    defer db.Unlock()
    
    // Truncate arrays back to batch start
    for tableName, startIdx := range batch.StartIndex {
        table := db.Tables[tableName]
        
        // Remove from index
        for i := startIdx; i < len(table.Items); i++ {
            key := table.Items[i].GetKey()
            delete(table.Index, key)
        }
        
        // Truncate array
        table.Items = table.Items[:startIdx]
        
        // Clear unresolved refs for rolled back items
        newUnresolved := []int{}
        for _, idx := range table.UnresolvedRefs {
            if idx < startIdx {
                newUnresolved = append(newUnresolved, idx)
            }
        }
        table.UnresolvedRefs = newUnresolved
    }
    
    batch.Status = BatchRolledBack
}
```

---

## The Linking Process (Your Multi-Pass Resolution)

### Reference Types and Resolution Order

```
┌─────────────────────────────────────────────────────────────┐
│                    REFERENCE TYPES                          │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  PASS 1 - Direct References (can resolve immediately):      │
│  ────────────────────────────────────────────────────────   │
│  ref      : Global lookup by key                            │
│             index["Table_" + key]                           │
│                                                             │
│  link     : Parent-scoped lookup                            │
│             index[parentIdx + "_Table_" + key]              │
│                                                             │
│  type_of  : Match child by type name                        │
│             parent.Childs.find(c => c.TypeName == value)    │
│                                                             │
│  up_copy  : Walk parent chain for ancestor type             │
│             current = parent; while current.Type != value   │
│                                                             │
│  PASS 2+ - Derived References (depend on prior resolution): │
│  ────────────────────────────────────────────────────────   │
│  ref_copy : Copy pointer from previous field                │
│             me.fieldP = me.prevFieldP.fieldP                │
│             Requires: prevField resolved                    │
│                                                             │
│  ref_child: Lookup in previous field's children             │
│             parent = me.prevFieldP                          │
│             me.fieldP = parent.index[key]                   │
│             Requires: prevField resolved                    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Resolution in Batch Context

```go
func (db *DatabaseT) ResolveBatchRefs(batch *BatchT) int {
    errs := 0
    
    // Process each table's unresolved items
    for tableName, table := range db.Tables {
        batchStart := batch.StartIndex[tableName]
        
        // Only process items in this batch
        for i := batchStart; i < len(table.Items); i++ {
            item := table.Items[i]
            
            // Pass 1: Direct references
            errs += db.resolveDirectRefs(item, i)
        }
    }
    
    // Pass 2+: Derived references (may need multiple passes)
    for pass := 0; pass < 5; pass++ {
        passErrs := 0
        
        for tableName, table := range db.Tables {
            batchStart := batch.StartIndex[tableName]
            
            for i := batchStart; i < len(table.Items); i++ {
                item := table.Items[i]
                passErrs += db.resolveDerivedRefs(item, i)
            }
        }
        
        if passErrs == 0 {
            break
        }
        errs = passErrs
    }
    
    return errs
}

func (db *DatabaseT) resolveDirectRefs(item Kp, idx int) int {
    errs := 0
    schema := item.Schema()
    
    for _, elem := range schema.Elements {
        if item.GetRefPointer(elem.Name) >= 0 {
            continue  // Already resolved
        }
        
        key := item.GetField(elem.Name).(string)
        
        switch elem.Type {
        case "ref":
            // Global lookup
            lookupKey := elem.TargetTable + "_" + key
            if targetIdx, ok := db.GlobalIndex[lookupKey]; ok {
                item.SetRefPointer(elem.Name, targetIdx)
            } else if elem.Check != "*" {
                errs++
            }
            
        case "link":
            // Parent-scoped lookup
            parentIdx := item.GetParentPointer()
            lookupKey := fmt.Sprintf("%d_%s_%s", parentIdx, elem.TargetTable, key)
            if targetIdx, ok := db.GlobalIndex[lookupKey]; ok {
                item.SetRefPointer(elem.Name, targetIdx)
            } else if elem.Check != "*" {
                errs++
            }
            
        case "type_of":
            // Match child by type
            for childIdx, child := range item.Children() {
                if child.TypeName() == key {
                    item.SetRefPointer(elem.Name, childIdx)
                    break
                }
            }
            
        case "up_copy":
            // Walk parent chain
            current := item.GetParentPointer()
            for current >= 0 {
                parent := db.Tables[item.ParentTable()].Items[current]
                if parent.TypeName() == key {
                    item.SetRefPointer(elem.Name, current)
                    break
                }
                current = parent.GetParentPointer()
            }
        }
    }
    
    return errs
}

func (db *DatabaseT) resolveDerivedRefs(item Kp, idx int) int {
    errs := 0
    schema := item.Schema()
    
    for i, elem := range schema.Elements {
        if item.GetRefPointer(elem.Name) >= 0 {
            continue  // Already resolved
        }
        
        switch elem.Type {
        case "ref_copy":
            // Copy from previous element's target
            if i == 0 {
                errs++
                continue
            }
            prevElem := schema.Elements[i-1]
            prevPtr := item.GetRefPointer(prevElem.Name)
            
            if prevPtr < 0 {
                errs++  // Previous not yet resolved
                continue
            }
            
            // Get the target of previous, then get ITS field
            prevTarget := db.Tables[prevElem.TargetTable].Items[prevPtr]
            item.SetRefPointer(elem.Name, prevTarget.GetRefPointer(elem.SourceField))
            
        case "ref_child":
            // Lookup in previous element's children
            if i == 0 {
                errs++
                continue
            }
            prevElem := schema.Elements[i-1]
            prevPtr := item.GetRefPointer(prevElem.Name)
            
            if prevPtr < 0 {
                errs++  // Previous not yet resolved
                continue
            }
            
            key := item.GetField(elem.Name).(string)
            prevTarget := db.Tables[prevElem.TargetTable].Items[prevPtr]
            
            // Look in previous target's children
            lookupKey := fmt.Sprintf("%d_%s_%s", prevPtr, elem.TargetTable, key)
            if targetIdx, ok := db.GlobalIndex[lookupKey]; ok {
                item.SetRefPointer(elem.Name, targetIdx)
            } else if elem.Check != "*" {
                errs++
            }
        }
    }
    
    return errs
}
```

---

## Continuous Operation Loop

```go
func (db *DatabaseT) Run(ctx context.Context, incoming <-chan []RawRecord) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
            
        case records := <-incoming:
            // Start new batch
            batch := db.StartBatch()
            
            // Append phase
            for _, rec := range records {
                item := db.parseRecord(rec)
                db.AppendBatch(batch, rec.Table, []Kp{item})
            }
            
            // Link phase
            errs := db.LinkBatch(batch)
            
            if errs > 0 {
                // Decide: partial commit or full rollback?
                if db.AllowPartialBatch {
                    db.MarkUnresolvedItems(batch)
                    db.CommitBatch(batch)
                } else {
                    db.RollbackBatch(batch)
                    db.ReportErrors(batch, errs)
                }
            } else {
                db.CommitBatch(batch)
            }
            
            // Trigger any watchers
            db.NotifyBatchComplete(batch)
        }
    }
}
```

---

## Integration with Actor System

### Session That Reads Database

```go
func (engine *EngineT) CreateQuerySession(output io.Writer) *SessionT {
    sess := &SessionT{
        ID:     uuid.New().String(),
        Schema: engine.Schema,
        Output: output,
        
        // Session gets READ access to database
        DB:     engine.Database,  // Shared reference
    }
    return sess
}

// Modified GetVar to read from database
func sGetVar(sess *SessionT, winp int, sc []string, va []string, lno string) (bool, string) {
    path := va
    
    // Check for database path prefix
    if len(path) > 0 && path[0] == "@" {
        return sess.DB.Query(path[1:], sc)
    }
    
    // Original behavior for session-local data
    rec := getPath(sess, winp, path, lno)
    // ...
}
```

### Actors That Write to Database

```go
case *KpAdd:
    target, path := parseAddPath(c.Kpath)
    
    switch target {
    case AddDatabase:
        // Must be in a batch context
        batch := sess.CurrentBatch
        if batch == nil {
            return fmt.Errorf("database write outside batch")
        }
        
        item := createItem(c, sess, winp)
        sess.Engine.Database.AppendBatch(batch, tableName, []Kp{item})
    }
```

### Complete Actor-Driven Batch

```
Actor ingest_batch .
  # Start transaction
  Transaction
  
  # Process each record in input
  This _.input_records process_record
  
  # Link all references
  Refs @
  
  # Validate
  All @.Component validate_refs
  
  # Commit if no errors
  Commit

Actor process_record Record
  # Transform and append to database
  Add @.Component {
    "name": "${name}",
    "type": "${type}",
    "parent": "${parent_key}"
  }

Actor validate_refs Component
  Its Element check_ref

Actor check_ref Element type in ref,link,ref_child
  Its comp check_resolved

Actor check_resolved Comp Kp ??
  # ref pointer is null
  C Error: ${.check_ref.parent.name}.${.check_ref.name} unresolved
  Break exit
```

---

## Storage Layout

```
┌─────────────────────────────────────────────────────────────┐
│                     STORAGE STRUCTURE                       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  data/                                                      │
│  ├── wal/                      # Write-Ahead Log            │
│  │   ├── 000001.wal           # Batch operations            │
│  │   ├── 000002.wal                                         │
│  │   └── current.wal          # Active log                  │
│  │                                                          │
│  ├── segments/                 # Frozen segments            │
│  │   ├── Component/                                         │
│  │   │   ├── seg_000001.dat   # Items 0-999                 │
│  │   │   ├── seg_000002.dat   # Items 1000-1999             │
│  │   │   └── seg_000003.dat   # Items 2000-2999             │
│  │   │                                                      │
│  │   ├── Element/                                           │
│  │   │   ├── seg_000001.dat                                 │
│  │   │   └── seg_000002.dat                                 │
│  │   │                                                      │
│  │   └── index/                # Persistent indices         │
│  │       ├── global.idx       # Find keys                   │
│  │       └── scoped.idx       # FindIn keys                 │
│  │                                                          │
│  ├── snapshots/                # Full state snapshots       │
│  │   └── 2024-01-15_000001/                                 │
│  │       ├── manifest.json                                  │
│  │       └── *.dat                                          │
│  │                                                          │
│  └── meta/                     # Metadata                   │
│      ├── schema.rio           # Current schema              │
│      ├── actors.rio           # Current actors              │
│      └── state.json           # Runtime state               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

### Segment Format

```go
type SegmentFile struct {
    Header SegmentHeader
    Items  []SerializedItem
    Index  []IndexEntry  // Offset table for random access
}

type SegmentHeader struct {
    Magic     [4]byte  // "RIOS"
    Version   uint32
    Table     string
    StartIdx  uint64
    EndIdx    uint64
    ItemCount uint64
    Created   int64
    Checksum  uint64
}

type SerializedItem struct {
    Length uint32
    Data   []byte  // MessagePack/Protobuf/JSON
}
```

---

## Summary: What Changes, What Stays

### Stays the Same
- Array-based storage (`ApComp`, `ApElement`)
- Index structures (`index["Comp_"+name]`)
- Multi-pass reference resolution
- Actor-based traversal and generation
- Win stack execution model

### Changes
| Component | Current | Continuous DB |
|-----------|---------|---------------|
| Arrays | In-memory only | Segmented, partially on disk |
| Lifetime | Single run | Persistent daemon |
| Batches | All-or-nothing | Transactional batches |
| Updates | Anywhere | Tail only |
| Resolution | Once at load | Per-batch + retry |
| Index | Rebuilt each run | Persistent + incremental |
| Output | stdout | Multiple sessions/streams |

### New Concepts
- **Batch**: Atomic unit of append + link
- **Tail**: Mutable frontier of each table
- **Segment**: Frozen, immutable historical chunk
- **WAL**: Durability for uncommitted batches
- **Session**: Isolated execution context for queries

This architecture preserves everything that works in the current system while adding the persistence and concurrency needed for a continuous database engine. The append-only with tail updates model is simpler than general-purpose databases and maps directly to how Rio already processes data.
