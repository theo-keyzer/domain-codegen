# Generated Python/JAX inference code for maxcut_qubo_model
# Problem: Max-Cut QUBO (N=100) using P-Bit Gibbs Sampling
# WARNING: Auto-generated file, do not edit manually

import jax
import jax.numpy as jnp
import numpy as np
from typing import Dict, Any, Tuple

# ... (pbit_probability, qubo_gibbs_update_kernel, calculate_qubo_energy remain unchanged) ...
@jax.jit
def pbit_probability(local_field: jnp.ndarray, beta: float) -> jnp.ndarray:
    """
    Calculates P(x_i=1) using the logistic (sigmoid) function from the local field (H_i).
    """
    return 1.0 / (1.0 + jnp.exp(-2.0 * beta * local_field))

# ... (qubo_gibbs_update_kernel remains unchanged) ...
@jax.jit
def qubo_gibbs_update_kernel(
    x_state: jnp.ndarray,    # Current state of binary variables (N=100)
    Q_matrix: jnp.ndarray,   # The N x N QUBO coupling matrix
    beta: float,             # Inverse Temperature (Annealing parameter)
    rng_key: jnp.ndarray     # JAX PRNG key for stochastic sampling
) -> Tuple[jnp.ndarray, jnp.ndarray]:
    N = Q_matrix.shape[0]
    local_fields = jnp.dot(Q_matrix, x_state)
    p_ones = pbit_probability(local_fields, beta)
    rng_key, subkey = jax.random.split(rng_key)
    random_numbers = jax.random.uniform(subkey, shape=(N,))
    new_x_state = jnp.where(random_numbers < p_ones, 1, 0).astype(x_state.dtype)
    return new_x_state, rng_key

# =========================================================================
# 2. MAX-CUT ENERGY AND VALUE CALCULATION (for verification)
# =========================================================================

@jax.jit
def calculate_qubo_energy(x_state: jnp.ndarray, Q_matrix: jnp.ndarray) -> float:
    """
    Calculates the QUBO energy (cost function value) for the current state.
    E = x^T * Q * x. Lower energy is the minimization goal.
    """
    energy = jnp.dot(x_state, jnp.dot(Q_matrix, x_state))
    return energy

@jax.jit
def calculate_max_cut_value(x_state: jnp.ndarray, Q_matrix: jnp.ndarray) -> float:
    """
    Calculates the actual Max-Cut value C(x) by summing the weights of all cut edges.
    """
    N = Q_matrix.shape[0]
    i_indices, j_indices = jnp.triu_indices(N, k=1)

    Q_values = Q_matrix[i_indices, j_indices]
    x_i_values = x_state[i_indices]
    x_j_values = x_state[j_indices]

    # Check where the cut occurs: cut_mask = 1 if x_i != x_j, else 0
    cut_mask = jnp.where(x_i_values != x_j_values, 1.0, 0.0)

    # Cut edge weight w_ij = -Q_ij (since Q_ij = -w_ij for Max-Cut QUBO)
    cut_contributions = -Q_values * cut_mask

    total_cut_value = jnp.sum(cut_contributions)

    return total_cut_value

# ... (run_maxcut_annealing remains unchanged) ...
def run_maxcut_annealing(
    Q_matrix: jnp.ndarray,
    anneal_steps: int,
    temp_start: float,
    temp_end: float,
    initial_state: jnp.ndarray,
    rng_key: jnp.ndarray
) -> Tuple[jnp.ndarray, float]:

    # ... (Annealing loop remains unchanged) ...

    log_temp_start = np.log(temp_start)
    log_temp_end = np.log(temp_end)

    temp_schedule = np.exp(np.linspace(log_temp_start, log_temp_end, anneal_steps))

    current_state = initial_state
    best_state = initial_state
    min_energy = calculate_qubo_energy(initial_state, Q_matrix)

    print(f"--- Starting P-Bit Max-Cut Annealing (N={Q_matrix.shape[0]}) ---")
    print(f"Anneal Steps: {anneal_steps}, T_start: {temp_start:.2f}, T_end: {temp_end:.2f}")

    for step in range(anneal_steps):
        temperature = temp_schedule[step]
        beta = 1.0 / temperature

        rng_key, update_key = jax.random.split(rng_key)

        new_state, rng_key = qubo_gibbs_update_kernel(
            current_state, Q_matrix, beta, update_key
        )

        current_state = new_state

        current_energy = calculate_qubo_energy(current_state, Q_matrix)

        if current_energy < min_energy:
            min_energy = current_energy
            best_state = current_state

        if step % (anneal_steps // 10) == 0 and step > 0:
            print(f"Step {step}/{anneal_steps} done. Current E: {current_energy:.4f}, Min E: {min_energy:.4f}")

    print(f"--- Annealing Complete ---")
    print(f"Final Minimum Energy Found: {min_energy:.4f}")
    return best_state, min_energy
# =========================================================================
# 4. DUMMY EXECUTION (Simulating the Max-Cut problem defined in pbit_maxcut.net)
# =========================================================================
if __name__ == "__main__":
    # --- Setup Parameters (from pbit_maxcut.net Config) ---
    N_VARS = 100
    ANNEAL_STEPS = 10000
    TEMP_START = 10.0
    TEMP_END = 0.01

    # --- Correct Dummy QUBO Matrix Construction for Max-Cut ---
    np.random.seed(42)
    Q_dummy = np.zeros((N_VARS, N_VARS), dtype=np.float32)

    # 1. Define positive edge weights (w_ij) and set Q_ij = -w_ij
    weights = np.zeros((N_VARS, N_VARS), dtype=np.float32)
    for i in range(N_VARS):
        for j in range(i + 1, N_VARS):
            if np.random.rand() < 0.1: # 10% sparsity
                w_ij = np.random.uniform(0.1, 1.0)
                weights[i, j] = weights[j, i] = w_ij
                Q_dummy[i, j] = Q_dummy[j, i] = -w_ij # Off-diagonal coupling

    # 2. Define Diagonal Biases (Q_ii) for Max-Cut: Q_ii = Sum_{j != i} w_ij
    # This term forces the QUBO minimization to maximize the cut.
    for i in range(N_VARS):
        Q_dummy[i, i] = np.sum(weights[i, :])

    Q_jax = jnp.array(Q_dummy)

    # --- Execution ---
    rng_key = jax.random.PRNGKey(1234)

    rng_key, init_key = jax.random.split(rng_key)
    initial_state = jax.random.randint(init_key, (N_VARS,), 0, 2, dtype=jnp.int32)

    final_state, min_energy = run_maxcut_annealing(
        Q_jax,
        ANNEAL_STEPS,
        TEMP_START,
        TEMP_END,
        initial_state,
        rng_key
    )

    # Calculate and print the actual Max-Cut value for the found minimum energy state
    max_cut_value = calculate_max_cut_value(final_state, Q_jax)

    print("\n--- Final Results ---")
    print(f"Best P-bit State (First 10): {final_state[:10].tolist()}...")
    print(f"Minimum QUBO Energy: {min_energy:.4f}")
    print(f"Actual Max-Cut Value: {max_cut_value:.4f}")
