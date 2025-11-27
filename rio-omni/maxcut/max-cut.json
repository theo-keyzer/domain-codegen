{
 "Agent": [
 ],
 "Objective": [
 ],
 "Memory": [
 ],
 "MemorySource": [
 ],
 "Thought": [
 ],
 "Citation": [
 ],
 "Action": [
 ],
 "ThoughtSource": [
 ],
 "Tool": [
 ],
 "Vector": [
 ],
 "Session": [
 ],
 "Snapshot": [
 ],
 "Restoration": [
 ],
 "CrossRef": [
 ],
 "Cycle": [
 ],
 "Recommendation": [
 ],
 "CanonCycle": [
 ],
 "Pattern": [
 ],
 "PatternSource": [
 ],
 "Evolution": [
 ],
 "MetricS": [
 ],
 "Artifact": [
 ],
 "GenExecution": [
 ],
 "GenValidation": [
 ],
 "GenLearning": [
 ],
 "GenDebug": [
 ],
 "Canon": [
 ],
 "Link": [
 ],
 "Section": [
 ],
 "CanonMeta": [
 ],
 "Implementation": [
 ],
 "CanonVersion": [
 ],
 "Library": [
 ],
 "GenProgram": [
 ],
 "GenBootstrap": [
 ],
 "GenInput": [
 ],
 "GenStrategy": [
 ],
 "GenSearchSpace": [
 ],
 "GenPatternRef": [
 ],
 "GenOutputSpec": [
 ],
 "GenMetric": [
 ],
 "GenEvolution": [
 ],
 "GenPipeline": [
 ],
 "GenStage": [
 ],
 "Project": [
  {
   "me": { "type": "Project", "id": 0 },
   "domain": { "type": "Domain", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "strategy": { "type": "Strategy", "id": 0 },
   "hardware": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "project",
     "_lno": "pbit-maxcut.omni:13",
     "desc": "Max-Cut QUBO problem on P-bit Memristor Array",
     "domain": "tsu",
     "hardware": "tsu_extropic_v1",
     "model": "maxcut_qubo_model",
     "project": "MaxCutPBit",
     "strategy": "simulated_annealing"
   }
  }
 ],
 "Domain": [
  {
   "me": { "type": "Domain", "id": 0 },
   "payload": {
     "_key": "name",
     "_lno": "pbit-maxcut.omni:8",
     "desc": "Thermodynamic Sampling Unit domain",
     "name": "tsu"
   }
  }
 ],
 "Hardware": [
  {
   "me": { "type": "Hardware", "id": 0 },
   "emulation": { "type": "Hardware", "id": 1 },
   "noise_model": { "type": "Constraint", "id": 0 },
   "payload": {
     "_key": "hardware",
     "_lno": "pbit-maxcut.omni:24",
     "backend": "tsu_analog",
     "emulation": "tsu_gpu_emulator",
     "hardware": "tsu_extropic_v1",
     "memory_gb": 16,
     "memristor_type": "TaOx",
     "noise_model": "thermal_noise_constraint",
     "operating_temp_k": 300,
     "pbit_count": 1024,
     "thermal_bandwidth_hz": 1000000000
   }
  },
  {
   "me": { "type": "Hardware", "id": 1 },
   "payload": {
     "_key": "hardware",
     "_lno": "pbit-maxcut.omni:36",
     "arch": "ampere",
     "backend": "cuda",
     "emulation_mode": "cycle_accurate",
     "hardware": "tsu_gpu_emulator",
     "memory_gb": 80
   }
  },
  {
   "me": { "type": "Hardware", "id": 2 },
   "payload": {
     "_key": "hardware",
     "_lno": "pbit-maxcut.omni:44",
     "backend": "cpu",
     "cores": 16,
     "desc": "Python/NumPy reference implementation",
     "hardware": "cpu_x86"
   }
  }
 ],
 "Model": [
  {
   "me": { "type": "Model", "id": 0 },
   "hardware": { "type": "Hardware", "id": 0 },
   "search_space": { "type": "SearchSpace", "id": 0 },
   "config": { "type": "Config", "id": 0 },
   "Layer": [0],
   "Tensor": [0, 1, 2, 3],
   "Op": [0],
   "payload": {
     "_key": "model",
     "_lno": "pbit-maxcut.omni:53",
     "code": "models/maxcut_qubo.py",
     "config": "annealing_config",
     "framework": "thrml_jax",
     "graph_structure": "dense",
     "hardware": "tsu_extropic_v1",
     "model": "maxcut_qubo_model",
     "num_variables": 100,
     "pgm_schema": "qubo",
     "search_space": "qubo_hyperparams",
     "type": "energy_based"
   }
  }
 ],
 "Layer": [
  {
   "me": { "type": "Layer", "id": 0 },
   "parent": { "type": "Model", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:80",
     "kParentp": "0",
     "layer": "qubo_sampling_layer",
     "layer_type": "sampling",
     "parent": "maxcut_qubo_model",
     "repeat": 1
   }
  }
 ],
 "Block": [
  {
   "me": { "type": "Block", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "payload": {
     "_key": "block",
     "_lno": "pbit-maxcut.omni:290",
     "block": "pbit_sampling_block",
     "block_type": "sampling",
     "coupling": "dense",
     "model": "maxcut_qubo_model",
     "noise_model": "thermal",
     "pbit_array_size": 100
   }
  }
 ],
 "Tensor": [
  {
   "me": { "type": "Tensor", "id": 0 },
   "parent": { "type": "Model", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:89",
     "desc": "Binary state vector for 100 P-bits",
     "dtype": "int8",
     "kParentp": "0",
     "parent": "maxcut_qubo_model",
     "role": "state",
     "shape": [
       100
     ],
     "tensor": "binary_state_vector",
     "value_range": [
       0,
       1
     ]
   }
  },
  {
   "me": { "type": "Tensor", "id": 1 },
   "parent": { "type": "Model", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:99",
     "desc": "QUBO coefficient matrix Q_ij",
     "dtype": "float32",
     "kParentp": "0",
     "layout": "symmetric",
     "parent": "maxcut_qubo_model",
     "role": "parameter",
     "shape": [
       100,
       100
     ],
     "tensor": "qubo_matrix"
   }
  },
  {
   "me": { "type": "Tensor", "id": 2 },
   "parent": { "type": "Model", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:109",
     "dtype": "float32",
     "kParentp": "0",
     "parent": "maxcut_qubo_model",
     "role": "input",
     "shape": [],
     "tensor": "temperature"
   }
  },
  {
   "me": { "type": "Tensor", "id": 3 },
   "parent": { "type": "Model", "id": 0 },
   "producer": { "type": "Op", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:117",
     "dtype": "float32",
     "kParentp": "0",
     "parent": "maxcut_qubo_model",
     "producer": "qubo_gibbs_update",
     "role": "output",
     "shape": [],
     "tensor": "energy_output"
   }
  }
 ],
 "Op": [
  {
   "me": { "type": "Op", "id": 0 },
   "parent": { "type": "Model", "id": 0 },
   "hardware": { "type": "Hardware", "id": 0 },
   "energy_fn": { "type": "EnergyFunction", "id": 0 },
   "search_space": { "type": "SearchSpace", "id": 0 },
   "strategy": { "type": "Strategy", "id": 0 },
   "layer": { "type": "Layer", "id": 0 },
   "Arg": [0, 1, 2, 3],
   "ControlFlow": [],
   "payload": {
     "_lno": "pbit-maxcut.omni:128",
     "calibration_mode": "adaptive",
     "code": "kernels/tsu/qubo_gibbs.cu",
     "desc": "Block Gibbs sampler for QUBO on TSU hardware",
     "energy_fn": "qubo_energy",
     "hardware": "tsu_extropic_v1",
     "kParentp": "0",
     "layer": "qubo_sampling_layer",
     "op": "qubo_gibbs_update",
     "op_type": "tsu_sampling",
     "parallel_flips": true,
     "parent": "maxcut_qubo_model",
     "search_space": "qubo_hyperparams",
     "strategy": "simulated_annealing",
     "sweeps_per_step": 10,
     "variant": "gibbs"
   }
  }
 ],
 "Arg": [
  {
   "me": { "type": "Arg", "id": 0 },
   "parent": { "type": "Op", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "tensor": { "type": "Tensor", "id": 1 },
   "payload": {
     "_lno": "pbit-maxcut.omni:147",
     "arg": "Q",
     "kParentp": "0",
     "parent": "qubo_gibbs_update",
     "role": "parameter",
     "tensor": "qubo_matrix"
   }
  },
  {
   "me": { "type": "Arg", "id": 1 },
   "parent": { "type": "Op", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "tensor": { "type": "Tensor", "id": 2 },
   "payload": {
     "_lno": "pbit-maxcut.omni:154",
     "arg": "T",
     "kParentp": "0",
     "parent": "qubo_gibbs_update",
     "role": "input",
     "tensor": "temperature"
   }
  },
  {
   "me": { "type": "Arg", "id": 2 },
   "parent": { "type": "Op", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "tensor": { "type": "Tensor", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:161",
     "arg": "state",
     "kParentp": "0",
     "parent": "qubo_gibbs_update",
     "role": "input_output",
     "tensor": "binary_state_vector"
   }
  },
  {
   "me": { "type": "Arg", "id": 3 },
   "parent": { "type": "Op", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "tensor": { "type": "Tensor", "id": 3 },
   "payload": {
     "_lno": "pbit-maxcut.omni:168",
     "arg": "energy",
     "kParentp": "0",
     "parent": "qubo_gibbs_update",
     "role": "output",
     "tensor": "energy_output"
   }
  }
 ],
 "Config": [
  {
   "me": { "type": "Config", "id": 0 },
   "schedule": { "type": "Strategy", "id": 0 },
   "payload": {
     "_key": "config",
     "_lno": "pbit-maxcut.omni:66",
     "anneal_steps": 10000,
     "batch_size": 128,
     "config": "annealing_config",
     "num_samples": 10000,
     "schedule": "simulated_annealing",
     "temp_end": 0.01,
     "temp_start": 10,
     "temperature_schedule": "geometric",
     "warmup_steps": 500
   }
  }
 ],
 "Kernel": [
 ],
 "EnergyFunction": [
  {
   "me": { "type": "EnergyFunction", "id": 0 },
   "payload": {
     "_key": "energy_fn",
     "_lno": "pbit-maxcut.omni:177",
     "code": "physics/qubo_energy.py",
     "coupling_type": "all_pairs",
     "energy_fn": "qubo_energy",
     "form": "E = sum_ij(Q_ij * x_i * x_j) + sum_i(Q_ii * x_i)"
   }
  }
 ],
 "SearchSpace": [
  {
   "me": { "type": "SearchSpace", "id": 0 },
   "target_model": { "type": "Model", "id": 0 },
   "Dimension": [0, 1, 2, 3],
   "payload": {
     "_key": "space",
     "_lno": "pbit-maxcut.omni:186",
     "algorithm": "bayesian",
     "budget": 100,
     "space": "qubo_hyperparams",
     "target_model": "maxcut_qubo_model"
   }
  }
 ],
 "Dimension": [
  {
   "me": { "type": "Dimension", "id": 0 },
   "parent": { "type": "SearchSpace", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:193",
     "dimension": "temperature_start",
     "kParentp": "0",
     "parent": "qubo_hyperparams",
     "range": [
       1,
       100
     ],
     "scale": "log",
     "type": "continuous"
   }
  },
  {
   "me": { "type": "Dimension", "id": 1 },
   "parent": { "type": "SearchSpace", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:201",
     "dimension": "temperature_end",
     "kParentp": "0",
     "parent": "qubo_hyperparams",
     "range": [
       0.001,
       1
     ],
     "scale": "log",
     "type": "continuous"
   }
  },
  {
   "me": { "type": "Dimension", "id": 2 },
   "parent": { "type": "SearchSpace", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:209",
     "dimension": "anneal_steps",
     "kParentp": "0",
     "parent": "qubo_hyperparams",
     "range": [
       1000,
       100000
     ],
     "scale": "log",
     "type": "integer"
   }
  },
  {
   "me": { "type": "Dimension", "id": 3 },
   "parent": { "type": "SearchSpace", "id": 0 },
   "payload": {
     "_lno": "pbit-maxcut.omni:217",
     "choices": [
       1,
       5,
       10,
       20,
       50
     ],
     "dimension": "sweeps_per_step",
     "kParentp": "0",
     "parent": "qubo_hyperparams",
     "type": "categorical"
   }
  }
 ],
 "Strategy": [
  {
   "me": { "type": "Strategy", "id": 0 },
   "search_space": { "type": "SearchSpace", "id": 0 },
   "fitness": { "type": "Metric", "id": 0 },
   "payload": {
     "_key": "strategy",
     "_lno": "pbit-maxcut.omni:226",
     "code": "strategies/simulated_annealing.py",
     "cooling": "geometric",
     "fitness": "cut_value_metric",
     "restarts": 5,
     "search_space": "qubo_hyperparams",
     "strategy": "simulated_annealing",
     "type": "annealing"
   }
  }
 ],
 "Constraint": [
  {
   "me": { "type": "Constraint", "id": 0 },
   "target_hw": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "constraint_id",
     "_lno": "pbit-maxcut.omni:254",
     "base_variance": 0.01,
     "constraint_id": "thermal_noise_constraint",
     "distribution": "gaussian",
     "target_hw": "tsu_extropic_v1",
     "temp_coefficient": 0.001,
     "type": "noise_model",
     "variance_model": "temperature_dependent"
   }
  },
  {
   "me": { "type": "Constraint", "id": 1 },
   "target_hw": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "constraint_id",
     "_lno": "pbit-maxcut.omni:264",
     "constraint_id": "power_budget",
     "max_watts": 50,
     "severity": "hard",
     "target_hw": "tsu_extropic_v1",
     "type": "power"
   }
  }
 ],
 "Metric": [
  {
   "me": { "type": "Metric", "id": 0 },
   "target_hw": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "metric",
     "_lno": "pbit-maxcut.omni:236",
     "desc": "Number of edges in the cut (maximize)",
     "metric": "cut_value_metric",
     "minimize": false,
     "target_hw": "tsu_extropic_v1",
     "unit": "cut_edges"
   }
  },
  {
   "me": { "type": "Metric", "id": 1 },
   "target_hw": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "metric",
     "_lno": "pbit-maxcut.omni:244",
     "desc": "QUBO energy value (minimize)",
     "metric": "energy_metric",
     "minimize": true,
     "target_hw": "tsu_extropic_v1",
     "unit": "energy"
   }
  }
 ],
 "Checkpoint": [
  {
   "me": { "type": "Checkpoint", "id": 0 },
   "model": { "type": "Model", "id": 0 },
   "payload": {
     "_key": "checkpoint_id",
     "_lno": "pbit-maxcut.omni:274",
     "checkpoint_id": "best_maxcut_solution",
     "metrics": "'''",
     "model": "maxcut_qubo_model",
     "state_path": "checkpoints/maxcut_best_20251121.pt",
     "timestamp": "2025-11-21T15:00:00Z"
   }
  }
 ],
 "Fusion": [
  {
   "me": { "type": "Fusion", "id": 0 },
   "hardware": { "type": "Hardware", "id": 0 },
   "payload": {
     "_key": "fusion",
     "_lno": "pbit-maxcut.omni:301",
     "fusion": "qubo_fused_update",
     "hardware": "tsu_extropic_v1",
     "ops_fused": [
       "gibbs_sweep",
       "energy_eval"
     ],
     "pattern_type": "sample_and_compute_energy",
     "speedup_factor": 2.3
   }
  }
 ],
 "ControlFlow": [
 ]
}
