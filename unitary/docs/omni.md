The Omni Architecture: New User Guide

1. The High-Level Workflow

The Omni system is a Graph Compilation Pipeline. It separates the "Intent" (what you want) from the "Implementation" (how it works). It is designed specifically to make AI Agents effective system architects.

The workflow transforms human/AI logic into executable code through four distinct stages:
Code snippet

graph LR
    A[AI / Human] -->|Writes| B(Rio Format .rio)
    B -->|Parsed by| C{Omni Linker}
    C -->|Validates via| D[Schema omni-doc]
    C -->|Compiles to| E(Omni JSON .json)
    E -->|Processed by| F[Actor System .act]
    F -->|Generates| G[Target Code .py/.cu/.v]

2. Phase I: Declaration (The Source Code)

File Format: .rio (Rio) Role: The Interface

The entry point is the Rio file. It is designed to be "AI-Friendly"—dense, token-efficient, and forgiving.

    Semantic Linking: You define relationships using names, not IDs (e.g., domain = tsu). The system resolves these connections later.

    Safe Code Injection: Uses ~~~ blocks to safely transport raw code (like Python or CUDA) without complex escaping, preventing syntax errors during AI generation.

    Flexible Data: The format allows you to invent new parameters on the fly. Any field not strictly defined as a link is treated as "Payload" and passed through to the generator automatically.

Example (.rio):
Code snippet

Project {
  project = MaxCutPBit
  domain = tsu          // Semantic link to another block
  voltage = 5.0         // Arbitrary payload data
}

3. Phase II: The Linker (The Middleware)

Definition: omni-doc.txt (Schema) Role: The Compiler

Between the Rio file and the JSON lies the Omni Linker. It uses the definitions in omni-doc.txt to stitch the loose graph together using a Multi-Pass Resolution algorithm.

    Pass 1 (Global Resolution): Locates all top-level components by unique key (e.g., finding the Hardware block named tsu_extropic_v1).

Pass 2 (Contextual Resolution): Resolves complex relative paths like up_copy (find my ancestor) or ref_child (find a sibling's child).

Structural vs. Payload: The schema strictly enforces the Graph Topology (how nodes connect) via types like ref and link , but allows the Data Payload to be flexible via the payload text field.

4. Phase III: Representation (The Intermediate Representation)

File Format: .json (Omni-JSON) Role: The Object File

The Linker outputs the JSON file. This is not for humans to write. It is a fully resolved "Object File" intended for machine consumption.

    Explicit Identity: Every component is assigned a unique integer ID (e.g., "id": 0).

Normalized Graph: Relationships are converted from names ("tsu") to explicit pointers (e.g., "domain": { "type": "Domain", "id": 0 }).

Payload Encapsulation: All flexible data defined in Rio (like memory_gb) is stored in a standardized payload block, keeping the main schema structure clean.

5. Phase IV: Synthesis (The Backend)

File Format: .act (Actor Script) Role: The Code Generator

The Actor System reads the JSON graph and generates the final source code. Unlike standard template engines (like Jinja2), the Actor system is a Graph Traversal Engine.

    Navigational Logic: Actors use commands like Its parent or Its model to walk the graph intelligently.

    Context Windows: The system maintains a "Window" of state. When an Actor runs on a Model, it inherently knows the context of that Model without manual variable passing.

    Modular Generation: Logic is broken into small units (e.g., gen_Layer, gen_Tensor) that can be reused and composed via the Du (Duplicate/Run) command.

Summary Reference

Component	File Ext	Analogy	User	Primary Goal
Rio	.rio	Source Code (.c)	AI / Human	Ease of Use: Low token cost, semantic naming.
Schema	.txt	Header Files (.h)	Architect	Integrity: Defines valid connections and resolution rules.
JSON	.json	Object File (.o)	System	Stability: Explicit IDs, fully resolved paths.
Actor	.act	Make/Linker	Developer	Synthesis: Graph traversal and file generation.
