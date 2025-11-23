The Omni Schema: Scope & Targets

1. Executive Summary

The Omni Schema is a unified definition language for Full-Stack Computing. Unlike standard ML formats (like ONNX) which only describe software graphs, or hardware descriptors (like Verilog) which only describe circuits, Omni models the interaction between the two.

Its primary purpose is to enable Physics-Aware Computing—systems where the noise, energy, and physical constraints of the hardware are intrinsic parts of the software algorithm.

2. Scope: The Four Pillars

The schema (omni-doc.txt) covers four distinct architectural areas that are traditionally siloed. By unifying them, the schema enables holistic system optimization.

A. The Compute Graph (Software)

    Components: Model, Layer, Tensor, Op, ControlFlow

    Scope: Defines the logical flow of data and operations.

    Key Difference: Unlike PyTorch/TensorFlow graphs, Omni graphs explicitly handle State (via Tensor roles) and Stochasticity (via EnergyFunction links), making them suitable for probabilistic computing (Bayesian networks, Boltzmann machines).

B. The Substrate (Hardware)

    Components: Hardware, Fusion, Kernel

    Scope: Describes the physical machine executing the graph.

    Key Features:

        Hierarchy: Racks → Nodes → Chips → Cores.

        Physics: Defines noise_model (e.g., thermal noise), power_budget, and bandwidth.

        Emulation: Allows high-level hardware (like a Quantum PU) to point to a digital twin emulator (GPU) for development.

C. The Physics (Thermodynamics)

    Components: EnergyFunction, Constraint, Checkpoint

    Scope: Bridges the gap between software logic and hardware reality.

    Usage: Used for Energy-Based Models (EBMs). Instead of defining a forward pass y = f(x), you define an energy landscape E(x) and let the hardware (or a solver) settle into the ground state.

D. The Meta-Optimizer (Search)

    Components: SearchSpace, Dimension, Strategy, Metric, Config

    Scope: Defines how the system should evolve or tune itself.

    Usage: Built-in Neural Architecture Search (NAS) and Hyperparameter Optimization (HPO). The schema creates a "Self-Optimizing Loop" where the Strategy modifies the Model to fit the Constraint.

3. Target Platforms

The Omni Schema is designed to compile Rio intents into executable code for three classes of hardware.

Class 1: Deterministic Digital (Standard)

    Targets: NVIDIA GPUs (CUDA), x86 CPUs, TPUs.

    Use Case: Standard Deep Learning, Simulation, Emulation.

    Mechanism: The schema compiles Op nodes into deterministic kernels (e.g., Matrix Multiplication).

Class 2: Stochastic Analog (Novel)

    Targets: Thermodynamic Sampling Units (TSUs), Memristor Arrays, Spintronics.

    Use Case: Combinatorial Optimization (Max-Cut, TSP), Generative AI Sampling, Bayesian Inference.

    Mechanism: The schema compiles EnergyFunction nodes directly into hardware control signals (e.g., voltage biases) to exploit physical noise for computation.

Class 3: Hybrid / Fused

    Targets: FPGA + CPU, GPU + Analog Co-processors.

    Use Case: Edge AI, Latency-Critical Control.

    Mechanism: The Fusion component maps specific sub-graphs to specific accelerators (e.g., "Run the Control Flow on CPU, but the Gibbs Sampling on TSU").

4. Architectural Stack Visualization

The following diagram illustrates how the Omni Schema connects these layers:

graph TD
    subgraph "Meta-Layer (Optimization)"
        S[Strategy] -->|Tunes| M[Model]
        S -->|Respects| C[Constraint]
    end

    subgraph "Logical Layer (Software)"
        M -->|Contains| Op[Operations]
        Op -->|Updates| T[Tensors]
        Op -->|Minimizes| E[Energy Function]
    end

    subgraph "Physical Layer (Hardware)"
        Op -->|Executes on| HW[Hardware]
        HW -->|Generates| N[Noise/Entropy]
        N -->|Drives| E
    end

5. Summary Matrix

Domain	Omni Components	Output Artifacts
Logic	Project, Domain, Model	Python (JAX/PyTorch)
Compute	Op, Kernel, Fusion	CUDA / OpenCL / Verilog
Data	Tensor, Checkpoint	.pt / .safetensors
Physics	EnergyFunction, Hardware	Simulation Configs / Analog Waveforms
Search	Strategy, Space, Metric	JSON Reports / Best Checkpoints

