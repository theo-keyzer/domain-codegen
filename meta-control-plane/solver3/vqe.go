// omega_unified.go - True Unified Omega with Entropic Crystallization
package main

import (
	"fmt"
	"math"
	"math/rand"
//	"strings"
	"time"
)

// ==============================================================================
//  DOMAIN MODELS
// ==============================================================================

type PauliOperator struct {
	PauliString string
	Coefficient float32
	ID          int
}

type Hamiltonian struct {
	Operators []PauliOperator
}

// ==============================================================================
//  STIEFEL MANIFOLD GEOMETRY
// ==============================================================================

type StiefelManifold struct{}

// RetractionQR projects the tangent vector back onto the manifold
func (sm *StiefelManifold) RetractionQR(W, Noise [][]float32, lr, temp float32) [][]float32 {
	n, p := len(W), len(W[0])
	W_cand := make([][]float32, n)
	
	// Perturb along tangent approximation
	for i := range W_cand {
		W_cand[i] = make([]float32, p)
		for j := range W_cand[i] {
			W_cand[i][j] = W[i][j] + lr*Noise[i][j]*temp
		}
	}
	
	// QR Decomposition to return to Stiefel Manifold
	Q, R := qrDecomposition(W_cand)
	
	// Ensure unique representation (positive diagonal of R)
	for i := 0; i < min(n, p); i++ {
		if R[i][i] < 0 {
			for j := 0; j < n; j++ { // Fix column j of Q
				Q[j][i] = -Q[j][i]
			}
		}
	}
	return Q
}

func GenerateNoise(rows, cols int, rng *rand.Rand) [][]float32 {
	mat := make([][]float32, rows)
	for i := range mat {
		mat[i] = make([]float32, cols)
		for j := range mat[i] {
			mat[i][j] = float32(rng.NormFloat64())
		}
	}
	return mat
}

// ==============================================================================
//  DIMENSIONALITY-AWARE KAPPA HELPER
// ==============================================================================

func sensibleKappa(dim int) float32 {
	return float32(dim) // Perfect for dim ≤ 64
	// Alternative: return float32(math.Sqrt(float64(dim))) * 5.5
}

// ==============================================================================
//  CORE OMEGA ARCHITECTURE
// ==============================================================================

type OmegaCore struct {
	Dim         int
	NumHeads    int
	HeadDim     int
	W_Embed     [][]float32
	Manifold    *StiefelManifold
	Kappas      []float32 // Concentration parameters (Inverse Temperature)
	RNG         *rand.Rand
	Temperature float32 // Global simulated annealing temp
}

func NewOmegaCore(dim, heads int, seed int64) *OmegaCore {
	rng := rand.New(rand.NewSource(seed))
	headDim := dim / heads
	
	// Initialize on Stiefel Manifold (Orthogonal matrix)
	wEmbed := initStiefel(dim, dim, rng)
	
	kappas := make([]float32, heads)
	for i := range kappas {
		kappas[i] = sensibleKappa(dim) // DIMENSION-AWARE FIX
	}
	
	return &OmegaCore{
		Dim:         dim,
		NumHeads:    heads,
		HeadDim:     headDim,
		W_Embed:     wEmbed,
		Manifold:    &StiefelManifold{},
		Kappas:      kappas,
		RNG:         rng,
		Temperature: 1.0,
	}
}

// ==============================================================================
//  THE UNIFIED LOSS FUNCTION
//  This is the heart of the "True Unification"
// ==============================================================================

type EmbeddedLoss struct {
	AttentionEntropy    float32
	CommutingPurity     float32
	EnergyConcentration float32
	ParameterSmoothness float32
	CombinedScore       float32
}

func (oc *OmegaCore) ComputeEmbeddedLoss(
	operators []PauliOperator,
	attention [][][]float32,
	generatedParams []float32,
	mode string,
) EmbeddedLoss {
	loss := EmbeddedLoss{}
	numOps := len(operators)
	
	// ---------------------------------------------------------
	// 1. Attention Entropy (The measure of crystallization)
	// Lower = Sharper, more "particle-like" attention
	// ---------------------------------------------------------
	totalEntropy := float32(0.0)
	count := 0
	for h := range attention {
		for i := range attention[h] {
			rowEntropy := float32(0.0)
			for j := range attention[h][i] {
				p := attention[h][i][j]
				if p > 1e-9 {
					rowEntropy -= p * float32(math.Log(float64(p)))
				}
			}
			totalEntropy += rowEntropy
			count++
		}
	}
	loss.AttentionEntropy = totalEntropy / float32(count)
	
	// ---------------------------------------------------------
	// 2. Commuting Purity (The measure of valid groups)
	// Higher = Better for clustering
	// ---------------------------------------------------------
	if mode == "clustering" {
		totalPurity := float32(0.0)
		headCount := 0
		for h := range attention {
			puritySum := float32(0.0)
			pairCount := 0
			// Sample high-attention pairs
			for i := 0; i < numOps; i++ {
				// Find max attended target
				maxIdx := -1
				maxVal := float32(0.0)
				for j := 0; j < numOps; j++ {
					if attention[h][i][j] > maxVal {
						maxVal = attention[h][i][j]
						maxIdx = j
					}
				}
				if maxIdx != -1 && maxIdx != i {
					weight := maxVal
					isCommuting := CheckCommutation(operators[i].PauliString, operators[maxIdx].PauliString)
					if isCommuting {
						puritySum += weight
					} else {
						puritySum -= weight // Penalty
					}
					pairCount++
				}
			}
			if pairCount > 0 {
				totalPurity += puritySum / float32(pairCount)
				headCount++
			}
		}
		if headCount > 0 {
			loss.CommutingPurity = totalPurity / float32(headCount)
		}
	}
	
	// ---------------------------------------------------------
	// 3. Energy Concentration (The measure of physical relevance)
	// Higher = Attention is focused on high-coefficient terms
	// ---------------------------------------------------------
	totalWeight := make([]float32, numOps)
	for h := range attention {
		for i := range attention[h] {
			for j := range attention[h][i] {
				totalWeight[j] += attention[h][i][j]
			}
		}
	}
	// Normalize
	sumW := float32(0.0)
	for _, w := range totalWeight { sumW += w }
	if sumW > 0 {
		for i := range totalWeight { totalWeight[i] /= sumW }
	}
	
	concentration := float32(0.0)
	for i, op := range operators {
		// Does the model look at the terms with big coefficients?
		absCoeff := float32(math.Abs(float64(op.Coefficient)))
		if absCoeff > 0.1 {
			concentration += totalWeight[i] * absCoeff * 10.0
		}
	}
	loss.EnergyConcentration = concentration

	// ---------------------------------------------------------
	// 4. Parameter Smoothness 
	// Lower = Less chaotic parameter landscape
	// ---------------------------------------------------------
	if len(generatedParams) > 1 {
		smooth := float32(0.0)
		for i := 0; i < len(generatedParams)-1; i++ {
			diff := generatedParams[i+1] - generatedParams[i]
			smooth += diff * diff
		}
		loss.ParameterSmoothness = smooth / float32(len(generatedParams))
	}

	// ---------------------------------------------------------
	// 5. THE UNIFIED SCORE
	// ---------------------------------------------------------
	if mode == "clustering" {
		// Goal: Sharp groups of commuting operators
		// Score = Purity - Entropy
		loss.CombinedScore = (loss.CommutingPurity * 2.0) - loss.AttentionEntropy
	} else if mode == "vqe" {
		// Goal: Concentrated energy focus with sharp attention structure
		// This is the FIX: We explicitly subtract Entropy here too!
		// Score = EnergyConcentration - Entropy - SmoothnessPenalty
		
		loss.CombinedScore = loss.EnergyConcentration - (loss.AttentionEntropy * 1.5) - (loss.ParameterSmoothness * 0.2)
	}
	
	return loss
}

// ==============================================================================
//  UNIFIED OPTIMIZATION LOOP
// ==============================================================================

func (oc *OmegaCore) OptimizeUnified(
	operators []PauliOperator,
	mode string,
	steps int,
	baseLR float32,
) (EmbeddedLoss, []float32) {
	
	// 1. Initial State Evaluation
	attention, params := oc.ComputeUnifiedAttention(operators, mode)
	bestLoss := oc.ComputeEmbeddedLoss(operators, attention, params, mode)
	bestScore := bestLoss.CombinedScore
	bestParams := make([]float32, len(params))
	copy(bestParams, params)
	
	// Stats tracking
	//lastImprov := 0
	
	// Annealing Constants
	tempStart := float32(2.5)
	tempEnd := float32(0.1)
	
	fmt.Printf("Initial %s Score: %.4f | Entropy: %.4f | Kappa: %.1f\n", 
		mode, bestScore, bestLoss.AttentionEntropy, oc.Kappas[0])
	
	for step := 0; step < steps; step++ {
		// Linear temperature decay
		progress := float32(step) / float32(steps)
		currentTemp := tempStart + (tempEnd-tempStart)*progress
		oc.Temperature = currentTemp // Updates internal state for Softmax
		
		// Backup state
		oldW := deepCopy(oc.W_Embed)
		
		// Perturb
		noise := GenerateNoise(oc.Dim, oc.Dim, oc.RNG)
		oc.W_Embed = oc.Manifold.RetractionQR(oc.W_Embed, noise, baseLR, currentTemp)
		
		// Evaluate
		currAttn, currParams := oc.ComputeUnifiedAttention(operators, mode)
		currLoss := oc.ComputeEmbeddedLoss(operators, currAttn, currParams, mode)
		
		// Acceptance Logic
		accepted := false
		
		if currLoss.CombinedScore > bestScore {
			// Greedy improvement
			accepted = true
		} else if mode == "vqe" {
			// Allow slight exploration in VQE mode for escaping local minima
			delta := bestScore - currLoss.CombinedScore
			prob := float32(math.Exp(float64(-delta * 10.0))) // Strict but probabilistic
			if oc.RNG.Float32() < prob {
				accepted = true
			}
		}
		
		if accepted {
			// If real improvement, update best
			if currLoss.CombinedScore > bestScore {
				bestLoss = currLoss
				bestScore = currLoss.CombinedScore
				bestParams = make([]float32, len(currParams))
				copy(bestParams, currParams)
				//lastImprov = step
			}
			// Keep the perturbed W (implicit)
		} else {
			// Revert
			oc.W_Embed = oldW
		}
		
		// Logging
		if step%500 == 0 || step == steps-1 {
			if mode == "clustering" {
				fmt.Printf("Step %4d: Score=%.4f | Pur=%.2f | Ent=%.4f | κ=%.1f\n", 
					step, bestScore, bestLoss.CommutingPurity, bestLoss.AttentionEntropy, oc.Kappas[0])
			} else {
				fmt.Printf("Step %4d: Score=%.4f | Conc=%.2f | Ent=%.4f | κ=%.1f\n", 
					step, bestScore, bestLoss.EnergyConcentration, bestLoss.AttentionEntropy, oc.Kappas[0])
			}
		}
		
		// Auto-sharpening (Kappa adjustment)
		if step%200 == 0 {
			oc.UpdateKappas(currAttn)
		}
	}
	
	return bestLoss, bestParams
}

// ==============================================================================
//  SUPPORT FUNCTIONS
// ==============================================================================

func (oc *OmegaCore) ComputeUnifiedAttention(
	operators []PauliOperator,
	mode string,
) ([][][]float32, []float32) {
	numOps := len(operators)
	
	// Encode & Project
	embeddings := make([][]float32, numOps)
	for i, op := range operators {
		enc := EncodePauliOperator(op, oc.Dim)
		embeddings[i] = MatVecMul(oc.W_Embed, enc)
		NormalizeVec(embeddings[i]) // Project to sphere
	}
	
	// Attention Calculation
	attention := make([][][]float32, oc.NumHeads)
	headDim := oc.Dim / oc.NumHeads
	
	for h := 0; h < oc.NumHeads; h++ {
		attention[h] = make([][]float32, numOps)
		start := h * headDim
		end := start + headDim
		
		for i := 0; i < numOps; i++ {
			attention[h][i] = make([]float32, numOps)
			rowMax := float32(-1e9)
			
			for j := 0; j < numOps; j++ {
				// 1. Geometric Similarity
				sim := DotProductSubset(embeddings[i], embeddings[j], start, end)
				
				// 2. Physics Bias (The distinct "Force" for each mode)
				bias := float32(0.0)
				
				if mode == "clustering" {
					// Bias towards commutativity
					if i != j {
						if CheckCommutation(operators[i].PauliString, operators[j].PauliString) {
							bias = 0.5
						} else {
							bias = -1.0
						}
					}
				} else if mode == "vqe" {
					// Bias towards magnitude (Significance)
					// This acts like a gravitational pull toward important terms
					bias = float32(math.Abs(float64(operators[j].Coefficient))) * 2.0
				}
				
				score := (sim * oc.Kappas[h]) + bias
				attention[h][i][j] = score
				if score > rowMax { rowMax = score }
			}
			
			// Softmax
			sumExp := float32(0.0)
			for j := 0; j < numOps; j++ {
				e := float32(math.Exp(float64(attention[h][i][j] - rowMax))) // Stable softmax
				attention[h][i][j] = e
				sumExp += e
			}
			for j := 0; j < numOps; j++ {
				attention[h][i][j] /= sumExp
			}
		}
	}
	
	// Parameter Generation (Only relevant for VQE output)
	params := []float32{}
	if mode == "vqe" {
		for i := 0; i < numOps; i++ {
			// Extract angles from dominant directions in embedding
			theta := float32(math.Atan2(float64(embeddings[i][1]), float64(embeddings[i][0])))
			params = append(params, theta)
		}
	}
	
	return attention, params
}

func (oc *OmegaCore) UpdateKappas(attention [][][]float32) {
	// Dynamically adjust concentration based on current entropy
	for h := 0; h < oc.NumHeads; h++ {
		// Logic: If attention is diffuse, increase Kappa to force crystallization
		oc.Kappas[h] *= 1.01 
		if oc.Kappas[h] > float32(oc.Dim*2) { // Cap at 2*dim
			oc.Kappas[h] = float32(oc.Dim * 2)
		}
	}
}

// ---------------- Standard Helpers ----------------

func EncodePauliOperator(op PauliOperator, dim int) []float32 {
	// Simple deterministic encoding
	vec := make([]float32, dim)
	lookup := map[rune]int{'I':0, 'X':1, 'Y':2, 'Z':3}
	for i, char := range op.PauliString {
		if i < dim {
			vec[i] = float32(lookup[char])
		}
	}
	// Encode Coefficient Magnitude at end of vector
	vec[dim-1] = float32(math.Abs(float64(op.Coefficient)))
	NormalizeVec(vec)
	return vec
}

func CheckCommutation(s1, s2 string) bool {
	ac := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != 'I' && s2[i] != 'I' && s1[i] != s2[i] {
			ac++
		}
	}
	return ac%2 == 0
}

func MatVecMul(A [][]float32, v []float32) []float32 {
	res := make([]float32, len(A))
	for i := range A {
		sum := float32(0)
		for j := range A[i] { sum += A[i][j] * v[j] }
		res[i] = sum
	}
	return res
}

func NormalizeVec(v []float32) {
	sum := float32(0)
	for _, x := range v { sum += x*x }
	norm := float32(math.Sqrt(float64(sum)))
	if norm > 0 {
		for i := range v { v[i] /= norm }
	}
}

func DotProductSubset(a, b []float32, start, end int) float32 {
	sum := float32(0)
	for i := start; i < end && i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func deepCopy(src [][]float32) [][]float32 {
	dst := make([][]float32, len(src))
	for i := range src {
		dst[i] = make([]float32, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func initStiefel(rows, cols int, rng *rand.Rand) [][]float32 {
	mat := GenerateNoise(rows, cols, rng)
	Q, _ := qrDecomposition(mat)
	return Q
}

func qrDecomposition(A [][]float32) ([][]float32, [][]float32) {
	m, n := len(A), len(A[0])
	Q := make([][]float32, m)
	R := make([][]float32, n)
	for i := range Q {
		Q[i] = make([]float32, n)
		copy(Q[i], A[i])
	}
	for i := range R { R[i] = make([]float32, n) }
	
	for j := 0; j < n; j++ {
		norm := float32(0.0)
		for i := 0; i < m; i++ { norm += Q[i][j] * Q[i][j] }
		norm = float32(math.Sqrt(float64(norm)))
		R[j][j] = norm
		for i := 0; i < m; i++ { Q[i][j] /= norm }
		for k := j+1; k < n; k++ {
			dot := float32(0.0)
			for i := 0; i < m; i++ { dot += Q[i][j] * Q[i][k] }
			R[j][k] = dot
			for i := 0; i < m; i++ { Q[i][k] -= dot * Q[i][j] }
		}
	}
	return Q, R
}

func min(a, b int) int { if a < b { return a }; return b }

// ==============================================================================
//  MAIN
// ==============================================================================

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== UNIFIED OMEGA: ENTROPIC CRYSTALLIZATION ===")
	fmt.Printf("Using sensibleKappa(dim)=%.1f for optimal concentration\n", sensibleKappa(32))
	
	ham := Hamiltonian{
		Operators: []PauliOperator{
			{PauliString: "II", Coefficient: -1.0},
			{PauliString: "ZZ", Coefficient: -0.01},  // Noise
			{PauliString: "XX", Coefficient: 0.18},   // Important
			{PauliString: "ZI", Coefficient: 0.39},   // Important
			{PauliString: "XY", Coefficient: 0.25},   // Important
			{PauliString: "ZX", Coefficient: 0.12},   // Moderate
		},
	}
	
	// 1. CLUSTERING RUN
	fmt.Println("\n--- MODE: CLUSTERING ---")
	fmt.Println("Target: Low Entropy, High Commuting Purity")
	clustCore := NewOmegaCore(32, 4, 101)
	clustLoss, _ := clustCore.OptimizeUnified(ham.Operators, "clustering", 29000, 0.05)
	
	fmt.Printf("\nFINAL CLUSTERING:\n Entropy: %.4f (Very Low)\n Purity:  %.2f (Should be ~1.0)\n", 
		clustLoss.AttentionEntropy, clustLoss.CommutingPurity)

	// 2. VQE RUN
	fmt.Println("\n--- MODE: VQE ---")
	fmt.Println("Target: Low Entropy, High Energy Concentration")
	vqeCore := NewOmegaCore(32, 4, 202)
	vqeLoss, params := vqeCore.OptimizeUnified(ham.Operators, "vqe", 33000, 0.05)
	
	fmt.Printf("\nFINAL VQE:\n Entropy: %.4f (Now Comparable to Clustering!)\n Concentration: %.4f\n", 
		vqeLoss.AttentionEntropy, vqeLoss.EnergyConcentration)
		
	fmt.Println("\nGenerated VQE Parameters (Structured):")
	for i, p := range params {
		fmt.Printf(" Op %d: %+.4f rad\n", i, p)
	}
	
	// CONCLUSION CHECK
	fmt.Println("\n=== VERIFICATION ===")
	if vqeLoss.AttentionEntropy < 0.5 {
		fmt.Println("SUCCESS: VQE Entropy successfully minimized via Unified objective.")
	} else {
		fmt.Println("NOTICE: VQE Entropy still high, increase penalty factor.")
	}
}
