package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

// ==============================================================================
//  DIMENSIONALITY-AWARE KAPPA HELPER
// ==============================================================================

func sensibleKappa(dim int) float32 {
	return float32(dim) // Perfect for dim ≤ 64
}

// ==============================================================================
//  PART 1: GEOMETRY & MANIFOLD
// ==============================================================================

type StiefelManifold struct{}

func (sm *StiefelManifold) RetractionQR(W, Noise [][]float32, lr, temp float32) [][]float32 {
	n, p := len(W), len(W[0])
	W_cand := make([][]float32, n)
	
	// Perturb
	for i := range W_cand {
		W_cand[i] = make([]float32, p)
		for j := range W_cand[i] {
			W_cand[i][j] = W[i][j] + lr * Noise[i][j] * temp
		}
	}
	
	// Project via QR
	Q, R := qrDecomposition(W_cand)
	
	// Sign Correction (Column Flip)
	for i := 0; i < min(n, p); i++ {
		if R[i][i] < 0 {
			for j := 0; j < n; j++ { Q[j][i] = -Q[j][i] }
		}
	}
	return Q
}

func GenerateNoise(rows, cols int, rng *rand.Rand) [][]float32 {
	mat := make([][]float32, rows)
	for i := range mat {
		mat[i] = make([]float32, cols)
		for j := range mat[i] { mat[i][j] = float32(rng.NormFloat64()) }
	}
	return mat
}

// ==============================================================================
//  PART 2: OMEGA LAYER (With Diagonal Penalty)
//  Fix: We ban i==j so the model MUST look elsewhere (Helix or Sheet)
// ==============================================================================

type OmegaLayer struct {
	Dim, NumHeads, HeadDim int
	LogKappa               float32
	WQ, WK                 [][]float32
	Manifold               *StiefelManifold
}

func NewOmegaLayer(dim, numHeads int, seed int64) *OmegaLayer {
	rng := rand.New(rand.NewSource(seed))
	kappa := sensibleKappa(dim)
	
	fmt.Printf("Initializing OmegaLayer: dim=%d, heads=%d, kappa=%.1f\n", 
		dim, numHeads, kappa)
	
	return &OmegaLayer{
		Dim:      dim,
		NumHeads: numHeads,
		HeadDim:  dim / numHeads,
		LogKappa: float32(math.Log(float64(kappa))), // DIMENSION-AWARE FIX
		WQ:       initStiefel(dim, dim, rng),
		WK:       initStiefel(dim, dim, rng),
		Manifold: &StiefelManifold{},
	}
}

func (l *OmegaLayer) ComputeAttention(x [][][]float32) [][][][]float32 {
	batch, seqLen := len(x), len(x[0])
	q := projectBatch(x, l.WQ)
	k := projectBatch(x, l.WK)
	qH := reshapeAndNorm(q, l.NumHeads, l.HeadDim)
	kH := reshapeAndNorm(k, l.NumHeads, l.HeadDim)
	
	kappa := float32(math.Exp(float64(l.LogKappa)))
	attn := make([][][][]float32, batch)
	
	for b := 0; b < batch; b++ {
		attn[b] = make([][][]float32, l.NumHeads)
		for h := 0; h < l.NumHeads; h++ {
			attn[b][h] = make([][]float32, seqLen)
			for i := 0; i < seqLen; i++ {
				attn[b][h][i] = make([]float32, seqLen)
				maxLogit := float32(-1e9)
				
				for j := 0; j < seqLen; j++ {
					// 1. Compute Geometric Score
					dot := float32(0.0)
					for d := 0; d < l.HeadDim; d++ {
						dot += qH[b][h][i][d] * kH[b][h][j][d]
					}
					val := dot * kappa
					
					// 2. CRITICAL FIX: SELF-ATTENTION PENALTY
					// Force the model to look at neighbors (Helix) or partners (Sheet)
					// by banning the exact diagonal and immediate neighbors.
					dist := math.Abs(float64(i - j))
					if dist < 2.0 {
						val -= 100.0 // Effective masking
					}

					attn[b][h][i][j] = val
					if val > maxLogit { maxLogit = val }
				}
				
				// Softmax
				sumExp := float32(0.0)
				for j := 0; j < seqLen; j++ {
					ex := float32(math.Exp(float64(attn[b][h][i][j] - maxLogit)))
					attn[b][h][i][j] = ex
					sumExp += ex
				}
				for j := 0; j < seqLen; j++ { attn[b][h][i][j] /= sumExp }
			}
		}
	}
	return attn
}

// ==============================================================================
//  PART 3: DATA GENERATION (Stronger Signal)
// ==============================================================================

func GenerateDiverseBioBatch(batchSize, length, dim int, rng *rand.Rand) ([][][]float32, []string) {
	data := make([][][]float32, batchSize)
	classLabels := make([]string, batchSize)
	
	for b := 0; b < batchSize; b++ {
		data[b] = make([][]float32, length)
		mode := "MIXED"
		if b % 3 == 0 { mode = "ALPHA_BUNDLE" }
		if b % 3 == 1 { mode = "BETA_BARREL" }
		classLabels[b] = mode
		
		// Fill with noise first
		for i := 0; i < length; i++ {
			data[b][i] = make([]float32, dim)
			for d := 0; d < dim; d++ { data[b][i][d] = float32(rng.NormFloat64()) }
		}

		// Alpha Helix: Strong local correlation (i interacts with i+4)
		if mode == "ALPHA_BUNDLE" || mode == "MIXED" {
			start, end := 4, length-4
			for i := start; i < end; i++ {
				for d := 0; d < dim; d++ {
					// 95% copy + 5% noise
					data[b][i][d] = 0.95*data[b][i-4][d] + 0.05*float32(rng.NormFloat64())
				}
			}
		}

		// Beta Sheet: Strong distant anti-diagonal correlation (i interacts with N-i)
		if mode == "BETA_BARREL" || mode == "MIXED" {
			for i := 0; i < length/2; i++ { // Only process first half
				pairIdx := length - 1 - i
				// 95% copy + 5% noise. This ensures the geometric vector is nearly identical.
				for d := 0; d < dim; d++ {
					data[b][pairIdx][d] = 0.95*data[b][i][d] + 0.05*float32(rng.NormFloat64())
				}
			}
		}
		
		// Normalize
		for i := 0; i < length; i++ {
			norm := float32(0.0)
			for d := 0; d < dim; d++ { norm += data[b][i][d] * data[b][i][d] }
			norm = float32(math.Sqrt(float64(norm))) + 1e-9
			for d := 0; d < dim; d++ { data[b][i][d] /= norm }
		}
	}
	return data, classLabels
}

// ==============================================================================
//  PART 4: ANALYSIS
// ==============================================================================

func AnalyzePhaseSeparation(attn [][][][]float32, labels []string) {
	heads := len(attn[0])
	seqLen := len(attn[0][0])
	
	fmt.Printf("\n%-4s | %-12s | %-12s | %-15s\n", "Head", "Helix (Loc)", "Sheet (Dist)", "Role")
	fmt.Println(strings.Repeat("-", 60))
	
	// Indices of specific class batches for verification
	alphaBatch := 0
	betaBatch := 1
	
	countHelix, countSheet := 0, 0
	
	for h := 0; h < heads; h++ {
		helixScore, sheetScore := 0.0, 0.0
		
		// Check Beta Batch for Sheet Patterns
		b := betaBatch
		totalBeta := 0.0
		for i := 0; i < seqLen; i++ {
			for j := 0; j < seqLen; j++ {
				p := float64(attn[b][h][i][j])
				totalBeta += p
				antiDist := math.Abs(float64((seqLen - 1 - j) - i))
				if antiDist < 4.0 && math.Abs(float64(i-j)) > 10 { // Strict anti-diagonal
					sheetScore += p
				}
			}
		}
		sheetScore /= totalBeta

		// Check Alpha Batch for Helix Patterns
		b = alphaBatch
		totalAlpha := 0.0
		for i := 0; i < seqLen; i++ {
			for j := 0; j < seqLen; j++ {
				p := float64(attn[b][h][i][j])
				totalAlpha += p
				dist := math.Abs(float64(i - j))
				if dist >= 3 && dist <= 5 { // Strict local (i-4)
					helixScore += p
				}
			}
		}
		helixScore /= totalAlpha
		
		role := "Unspecialized"
		// Threshold lowered slightly as diagonal penalty disperses some attention
		if helixScore > 0.25 && helixScore > sheetScore {
			role = "HELIX (🧬)"
			countHelix++
		} else if sheetScore > 0.25 && sheetScore > helixScore {
			role = "SHEET (📄)"
			countSheet++
		} 
		
		fmt.Printf(" #%d  |   %.3f      |   %.3f      | %s\n", 
			h, helixScore, sheetScore, role)
			
		// Print map for first found sheet head
		if role == "SHEET (📄)" && countSheet == 1 {
			PrintHeatmap(attn[betaBatch][h], h, "Beta Sheet Detection (Anti-Diagonal)")
		}
	}
	
	if countHelix > 0 && countSheet > 0 {
		fmt.Println("\nSUCCESS: Bifurcation Achieved. Heads specialized into distinct physical regimes.")
	} else {
		fmt.Println("\nPENDING: Try increasing batch size or steps.")
	}
}

func PrintHeatmap(attn [][]float32, headID int, title string) {
	fmt.Printf("\n--- %s ---\n", title)
	size := len(attn)
	scale := 2
	for i := 0; i < size; i += scale {
		if i%8==0 { fmt.Printf("%2d ", i) } else { fmt.Print("   ") }
		for j := 0; j < size; j += scale {
			val := attn[i][j]
			char := "·"
			if val > 0.4 { char = "×" } // Sheet
			if val > 0.7 { char = "█" } // Strong
			fmt.Print(char)
		}
		fmt.Println()
	}
}

// ==============================================================================
//  PART 5: MAIN & UTILS
// ==============================================================================

func main() {
	seed := int64(101) // Changed seed for fresh randomness
	rng := rand.New(rand.NewSource(seed))
	
	fmt.Println("############################################################")
	fmt.Println("# OMEGA: PHASE SEPARATION (Step 3: Diagonal Masking)       #")
	fmt.Println("# WITH DIMENSION-AWARE KAPPA (κ ≈ dim = 64)                #")
	fmt.Println("############################################################")
	
	dim, heads, seqLen := 64, 8, 48
	data, classes := GenerateDiverseBioBatch(12, seqLen, dim, rng)
	omega := NewOmegaLayer(dim, heads, seed)
	
	steps := 9000 // Faster convergence expected with correct kappa
	baseLr := float32(0.18)
	bestEnt := float32(100.0)
	
	fmt.Printf("Starting with kappa=%.1f (exp(LogKappa))\n", math.Exp(float64(omega.LogKappa)))
	
	for i := 0; i < steps; i++ {
		progress := float32(i) / float32(steps)
		temp := 2.5 * (1.0 - progress) + 0.05
		
		oldWQ, oldWK := deepCopy(omega.WQ), deepCopy(omega.WK)
		nQ, nK := GenerateNoise(dim, dim, rng), GenerateNoise(dim, dim, rng)
		
		omega.WQ = omega.Manifold.RetractionQR(omega.WQ, nQ, baseLr, temp)
		omega.WK = omega.Manifold.RetractionQR(omega.WK, nK, baseLr, temp)
		
		newAttn := omega.ComputeAttention(data)
		h := ComputeEntropy(newAttn)
		
		if h < bestEnt {
			bestEnt = h
			if i%3 == 0 { fmt.Printf("Step %4d | T %.2f | Ent %.4f [New Best]\n", i, temp, h) }
		} else {
			// Metropolis acceptance
			delta := h - bestEnt
			if rng.Float64() < math.Exp(float64(-delta*15.0/float32(temp))) {
				// Accept degradation
			} else {
				omega.WQ = oldWQ; omega.WK = oldWK
			}
		}
	}
	
	fmt.Printf("\nFinal Entropy: %.4f\n", bestEnt)
	AnalyzePhaseSeparation(omega.ComputeAttention(data), classes)
}

// Standard Utils...
func ComputeEntropy(attn [][][][]float32) float32 {
	total, count := float32(0.0), 0
	for b := range attn { for h := range attn[b] { for i := range attn[b][h] {
				e := float32(0.0)
				for j := range attn[b][h][i] { p := attn[b][h][i][j]
					if p > 1e-9 { e -= p * float32(math.Log(float64(p))) }}
				total += e; count++ }}}
	return total / float32(count)
}
func deepCopy(src [][]float32) [][]float32 { dst := make([][]float32, len(src)); for i := range src { dst[i] = make([]float32, len(src[i])); copy(dst[i], src[i]) }; return dst }
func initStiefel(rows, cols int, rng *rand.Rand) [][]float32 { mat := GenerateNoise(rows, cols, rng); Q, _ := qrDecomposition(mat); return Q }
func qrDecomposition(A [][]float32) ([][]float32, [][]float32) { m, n := len(A), len(A[0]); Q := make([][]float32, m); R := make([][]float32, n); for i := range Q { Q[i] = make([]float32, n); copy(Q[i], A[i]) }; for i := range R { R[i] = make([]float32, n) }; for j := 0; j < n; j++ { norm := float32(0.0); for i := 0; i < m; i++ { norm += Q[i][j] * Q[i][j] }; norm = float32(math.Sqrt(float64(norm))); R[j][j] = norm; for i := 0; i < m; i++ { Q[i][j] /= norm }; for k := j + 1; k < n; k++ { dot := float32(0.0); for i := 0; i < m; i++ { dot += Q[i][j] * Q[i][k] }; R[j][k] = dot; for i := 0; i < m; i++ { Q[i][k] -= dot * Q[i][j] }}}; return Q, R }
func projectBatch(x [][][]float32, W [][]float32) [][][]float32 { b, s, d := len(x), len(x[0]), len(W); out := make([][][]float32, b); for i:=0; i<b; i++ { out[i] = make([][]float32, s); for j:=0; j<s; j++ { out[i][j] = make([]float32, d); for k:=0; k<d; k++ { sum := float32(0.0); for l:=0; l<len(W[0]); l++ { sum += x[i][j][l]*W[k][l] }; out[i][j][k] = sum }}}; return out }
func reshapeAndNorm(x [][][]float32, heads, headDim int) [][][][]float32 { b, s := len(x), len(x[0]); out := make([][][][]float32, b); for i := 0; i < b; i++ { out[i] = make([][][]float32, heads); for h := 0; h < heads; h++ { out[i][h] = make([][]float32, s); for j := 0; j < s; j++ { out[i][h][j] = make([]float32, headDim); norm := float32(0.0); for d := 0; d < headDim; d++ { v := x[i][j][h*headDim+d]; out[i][h][j][d] = v; norm += v*v }; norm = float32(math.Sqrt(float64(norm))) + 1e-9; for d := 0; d < headDim; d++ { out[i][h][j][d] /= norm }}}}; return out }
func min(a, b int) int { if a < b { return a }; return b }
