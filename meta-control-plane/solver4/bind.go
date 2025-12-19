package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)
// https://www.pgnmentor.com/files.html#openings
// ====================== COMMON UTILITIES ======================

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func deepCopy(src [][]float32) [][]float32 {
	dst := make([][]float32, len(src))
	for i := range src {
		dst[i] = make([]float32, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
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

func NormalizeVec(v []float32) {
	sum := float32(0)
	for _, x := range v {
		sum += x * x
	}
	norm := float32(math.Sqrt(float64(sum))) + 1e-9
	for i := range v {
		v[i] /= norm
	}
}

func MatVecMul(A [][]float32, v []float32) []float32 {
	res := make([]float32, len(A))
	for i := range A {
		sum := float32(0)
		for j := range A[i] {
			sum += A[i][j] * v[j]
		}
		res[i] = sum
	}
	return res
}

func DotProductSubset(a, b []float32, start, end int) float32 {
	sum := float32(0)
	for i := start; i < end && i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func qrDecomposition(A [][]float32) ([][]float32, [][]float32) {
	m, n := len(A), len(A[0])
	Q := make([][]float32, m)
	for i := range Q {
		Q[i] = make([]float32, n)
		copy(Q[i], A[i])
	}
	R := make([][]float32, n)
	for i := range R {
		R[i] = make([]float32, n)
	}

	for j := 0; j < n; j++ {
		norm := float32(0.0)
		for i := 0; i < m; i++ {
			norm += Q[i][j] * Q[i][j]
		}
		norm = float32(math.Sqrt(float64(norm))) + 1e-9
		R[j][j] = norm
		for i := 0; i < m; i++ {
			Q[i][j] /= norm
		}
		for k := j + 1; k < n; k++ {
			dot := float32(0.0)
			for i := 0; i < m; i++ {
				dot += Q[i][j] * Q[i][k]
			}
			R[j][k] = dot
			for i := 0; i < m; i++ {
				Q[i][k] -= dot * Q[i][j]
			}
		}
	}
	return Q, R
}

func initStiefel(rows, cols int, rng *rand.Rand) [][]float32 {
	mat := GenerateNoise(rows, cols, rng)
	Q, _ := qrDecomposition(mat)
	return Q
}

type StiefelManifold struct{}

func (sm *StiefelManifold) RetractionQR(W, Noise [][]float32, lr, temp float32) [][]float32 {
	n, p := len(W), len(W[0])
	W_cand := make([][]float32, n)
	for i := range W_cand {
		W_cand[i] = make([]float32, p)
		for j := range W_cand[i] {
			W_cand[i][j] = W[i][j] + lr*Noise[i][j]*temp
		}
	}
	Q, R := qrDecomposition(W_cand)
	for i := 0; i < min(n, p); i++ {
		if R[i][i] < 0 {
			for j := 0; j < n; j++ {
				Q[j][i] = -Q[j][i]
			}
		}
	}
	return Q
}

func sensibleKappa(dim int) float32 {
	return float32( math.Sqrt(float64(dim)) * 8.0 )
	//return( float32(56.0) )
	//return float32(dim)
}

// ====================== OMEGA CORE ======================

type OmegaCore struct {
	Dim      int
	NumHeads int
	HeadDim  int
	W_Embed  [][]float32
	Manifold *StiefelManifold
	Kappas   []float32
	RNG      *rand.Rand
}

func NewOmegaCore(dim, heads int, seed int64) *OmegaCore {
	rng := rand.New(rand.NewSource(seed))
	headDim := dim / heads
	w := initStiefel(dim, dim, rng)
	kappas := make([]float32, heads)
	for i := range kappas {
		kappas[i] = sensibleKappa(dim)
	}
	return &OmegaCore{
		Dim:      dim,
		NumHeads: heads,
		HeadDim:  headDim,
		W_Embed:  w,
		Manifold: &StiefelManifold{},
		Kappas:   kappas,
		RNG:      rng,
	}
}

// General attention (works for both domains)
func (oc *OmegaCore) ComputeAttention(inputs [][][]float32, diagonalPenalty bool, extraBias func(b, i, j int) float32) [][][][]float32 {
	batch := len(inputs)
	if batch == 0 {
		return nil
	}
	//seqLens := len(inputs[0])
	attn := make([][][][]float32, batch)

	for b := 0; b < batch; b++ {
		seqLen := len(inputs[b])
		embed := make([][]float32, seqLen)
		for i := 0; i < seqLen; i++ {
			embed[i] = MatVecMul(oc.W_Embed, inputs[b][i])
			NormalizeVec(embed[i])
		}

		attn[b] = make([][][]float32, oc.NumHeads)
		for h := 0; h < oc.NumHeads; h++ {
			attn[b][h] = make([][]float32, seqLen)
			start := h * oc.HeadDim
			end := start + oc.HeadDim
			for i := 0; i < seqLen; i++ {
				attn[b][h][i] = make([]float32, seqLen)
				maxL := float32(-1e9)
				for j := 0; j < seqLen; j++ {
					sim := DotProductSubset(embed[i], embed[j], start, end)
					val := sim * oc.Kappas[h]
					if diagonalPenalty && math.Abs(float64(i-j)) < 2.0 {
						val -= 100.0
					}
					if extraBias != nil {
						val += extraBias(b, i, j)
					}
					attn[b][h][i][j] = val
					if val > maxL {
						maxL = val
					}
				}
				// stable softmax
				sumExp := float32(0)
				for j := 0; j < seqLen; j++ {
					e := float32(math.Exp(float64(attn[b][h][i][j] - maxL)))
					attn[b][h][i][j] = e
					sumExp += e
				}
				if sumExp > 0 {
					for j := 0; j < seqLen; j++ {
						attn[b][h][i][j] /= sumExp
					}
				}
			}
		}
	}
	return attn
}

func ComputeEntropy(attn [][][][]float32) float32 {
	total := float32(0)
	count := 0
	for b := range attn {
		for h := range attn[b] {
			for i := range attn[b][h] {
				e := float32(0)
				for j := range attn[b][h][i] {
					p := attn[b][h][i][j]
					if p > 1e-9 {
						e -= p * float32(math.Log(float64(p)))
					}
				}
				total += e
				count++
			}
		}
	}
	if count > 0 {
		return total / float32(count)
	}
	return 0
}

// ====================== PROTEIN DATA ======================

func GenerateDiverseBioBatch(batchSize, length, dim int, rng *rand.Rand) ([][][]float32, []string) {
	data := make([][][]float32, batchSize)
	labels := make([]string, batchSize)
	for b := 0; b < batchSize; b++ {
		data[b] = make([][]float32, length)
		mode := "MIXED"
		if b%3 == 0 {
			mode = "ALPHA_BUNDLE"
		}
		if b%3 == 1 {
			mode = "BETA_BARREL"
		}
		labels[b] = mode

		for i := 0; i < length; i++ {
			data[b][i] = make([]float32, dim)
			for d := 0; d < dim; d++ {
				data[b][i][d] = float32(rng.NormFloat64())
			}
		}

		if mode == "ALPHA_BUNDLE" || mode == "MIXED" {
			for i := 4; i < length-4; i++ {
				for d := 0; d < dim; d++ {
					data[b][i][d] = 0.95*data[b][i-4][d] + 0.05*float32(rng.NormFloat64())
				}
			}
		}

		if mode == "BETA_BARREL" || mode == "MIXED" {
			for i := 0; i < length/2; i++ {
				pair := length - 1 - i
				for d := 0; d < dim; d++ {
					data[b][pair][d] = 0.95*data[b][i][d] + 0.05*float32(rng.NormFloat64())
				}
			}
		}

		for i := 0; i < length; i++ {
			NormalizeVec(data[b][i])
		}
	}
	return data, labels
}

// ====================== QUANTUM DATA ======================

type PauliOperator struct {
	PauliString string
	Coefficient float32
}

func EncodePauliOperator(op PauliOperator, dim int) []float32 {
	vec := make([]float32, dim)
	lookup := map[rune]int{'I': 0, 'X': 1, 'Y': 2, 'Z': 3}
	n := len(op.PauliString)
	for i := 0; i < n && i < dim; i++ {
		vec[i] = float32(lookup[rune(op.PauliString[i])])
	}
	vec[dim-1] = float32(math.Abs(float64(op.Coefficient)))
	NormalizeVec(vec)
	return vec
}

func getQuantumBatches(dim int) ([][][]float32, [][]PauliOperator) {
	// Two simple Hamiltonians as separate "batch" items
	hams := [][]PauliOperator{
		{
			{"II", -1.0},
			{"ZZ", 0.1},
			{"XX", 0.18},
			{"ZI", 0.39},
			{"XY", 0.25},
		},
		{
			{"III", -0.8},
			{"XXX", 0.3},
			{"YYY", 0.2},
			{"ZIZ", 0.4},
		},
	}

	batch := make([][][]float32, len(hams))
	allOps := make([][]PauliOperator, len(hams))
	for b, ham := range hams {
		allOps[b] = ham
		n := len(ham)
		batch[b] = make([][]float32, n)
		for i, op := range ham {
			batch[b][i] = EncodePauliOperator(op, dim)
		}
	}
	return batch, allOps
}

func Predict8StateSecondaryStructure(core *OmegaCore, sequence [][]float32) []string {
    attn := core.ComputeAttention([][][]float32{sequence}, true, nil)[0]
    heads := len(attn)
    seqLen := len(sequence)
    
    // DSSP-like 8 states: H=alpha-helix, G=310-helix, I=pi-helix, E=beta-sheet, B=beta-bridge, T=turn, S=bend, C=coil
    predictions := make([]string, seqLen)
    
    for i := 0; i < seqLen; i++ {
        // Collect attention statistics
        localPattern := make([]float32, 8) // 0-7 distance bins
        for h := 0; h < heads; h++ {
            for j := 0; j < seqLen; j++ {
                dist := int(math.Abs(float64(i - j)))
                if dist < 8 {
                    localPattern[dist] += attn[h][i][j]
                }
            }
        }
        
        // Heuristic rules (your model already knows these!)
        if localPattern[4] > 0.3 && localPattern[4] > localPattern[3] && localPattern[4] > localPattern[5] {
            predictions[i] = "H" // Classic alpha-helix (i→i+4)
        } else if localPattern[3] > 0.25 && localPattern[3] > localPattern[4] {
            predictions[i] = "G" // 310-helix (i→i+3)
        } else if localPattern[5] > 0.3 {
            predictions[i] = "I" // Pi-helix (i→i+5)
        } else if i > 0 && i < seqLen-1 {
            // Check for beta patterns (anti-parallel)
            antiPartner := seqLen - 1 - i
            betaScore := float32(0)
            for h := 0; h < heads; h++ {
                betaScore += attn[h][i][antiPartner]
            }
            if betaScore > 0.4 {
                predictions[i] = "E" // Beta-sheet
            } else if betaScore > 0.2 {
                predictions[i] = "B" // Beta-bridge
            } else if localPattern[1]+localPattern[2] > 0.5 {
                predictions[i] = "T" // Turn
            } else if localPattern[6]+localPattern[7] > 0.4 {
                predictions[i] = "S" // Bend
            } else {
                predictions[i] = "C" // Coil
            }
        } else {
            predictions[i] = "C"
        }
    }
    
    // Post-process: extend helices and sheets
    //for i := 1; i < seqLen-1; i++ {
    //    if predictions[i] == "H" && (predictions[i-1] == "H" || predictions[i+1] == "H") {
    //        predictions[i] = "H"
    //    }
    //}
    
    return predictions
}

// ====================== MAIN: JOINT TRAINING ======================

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	dim := 128
	heads := 32 // more heads to allow rich specialization
	core := NewOmegaCore(dim, heads, 123)

	fmt.Println("=== UNIFIED OMEGA: JOINT PROTEIN + QUANTUM TRAINING ===")

	// Fixed datasets (can regenerate periodically if desired)
	proteinData, proteinLabels := GenerateDiverseBioBatch(8, 48, dim, rng)
	quantumData, quantumOps := getQuantumBatches(dim)

	steps := 20000 // 50000
	baseLr := float32(0.12)
	bestScore := float32(-9999)
	bestW := deepCopy(core.W_Embed)

	for step := 0; step < steps; step++ {
		progress := float32(step) / float32(steps)
		//temp := 2.8*(1-progress) + 0.05
		temp := 1.4*(1-progress) + 0.02

		oldW := deepCopy(core.W_Embed)
		noise := GenerateNoise(dim, dim, core.RNG)
		core.W_Embed = core.Manifold.RetractionQR(core.W_Embed, noise, baseLr, temp)

		// Protein side
    helixBias := func(b, i, j int) float32 {
        if proteinLabels[b] == "ALPHA_BUNDLE" && math.Abs(float64(i-j)) >= 3 && math.Abs(float64(i-j)) <= 5 {
            return 1.5 // 0.8  // Strong bonus for helix patterns in alpha bundles
        }
        return 0
    }
		protAttn := core.ComputeAttention(proteinData, true, helixBias)
		entProt := ComputeEntropy(protAttn)

		// Quantum side - example bias: +coeff magnitude for important terms, +0.5 if commuting
//	if step%20 == 0 {
		quantBias := func(b, i, j int) float32 {
			if i == j {
				return 0
			}
			opI := quantumOps[b][i]
			opJ := quantumOps[b][j]
			bias := float32(math.Abs(float64(opJ.Coefficient))) * 1.5
			if CheckCommutation(opI.PauliString, opJ.PauliString) {
				bias += 0.5
			} else {
				bias -= 0.3
			}
			return bias
		}
		quantAttn := core.ComputeAttention(quantumData, false, quantBias)
		entQuant := ComputeEntropy(quantAttn)
//	}
		// Bonus for quantum concentration (simple proxy)
		// Higher attention on high-coeff terms
		concentrationBonus := float32(0)
		for b := range quantAttn {
			for h := range quantAttn[b] {
				for i := range quantAttn[b][h] {
					for j := range quantAttn[b][h][i] {
						concentrationBonus += quantAttn[b][h][i][j] * float32(math.Abs(float64(quantumOps[b][j].Coefficient)))
					}
				}
			}
		}

		score := - (entProt + entQuant) + concentrationBonus*0.5

		if score > bestScore {
			bestScore = score
			bestW = deepCopy(core.W_Embed)
			fmt.Printf("Step %5d | EntP %.3f | EntQ %.3f | Bonus %.3f | Score %.3f ★\n",
				step, entProt, entQuant, concentrationBonus, score)
		} else {
			delta := score - bestScore
			if core.RNG.Float64() < math.Exp(float64(delta/temp*10)) {
				// accept worse
			} else {
				core.W_Embed = oldW
			}
		}

		// Gentle kappa sharpening
		if step%300 == 0 {
			for h := range core.Kappas {
				core.Kappas[h] *= 1.008
				if core.Kappas[h] > float32(dim*3) {
					core.Kappas[h] = float32(dim * 3)
				}
			}
		}
	}

	core.W_Embed = bestW
	fmt.Printf("\nJoint training finished. Best score: %.4f\n", bestScore)

	// ====================== QUERY PROTEIN ANGLE ======================
	fmt.Println("\n=== QUERY: Protein Folding Specialization ===")
	protAttnFinal := core.ComputeAttention(proteinData, true, nil)
	fmt.Printf("Final protein entropy: %.4f\n", ComputeEntropy(protAttnFinal))
	// Plug in your AnalyzePhaseSeparation / PrintHeatmap here for full visualization

	// ====================== QUERY QUANTUM ANGLE ======================
	fmt.Println("\n=== QUERY: Quantum Clustering / Energy Focus ===")
	quantAttnFinal := core.ComputeAttention(quantumData, false, nil) // or with bias for params
	fmt.Printf("Final quantum entropy: %.4f\n", ComputeEntropy(quantAttnFinal))
	// Add purity/concentration analysis or parameter extraction here

	fmt.Println("\nModel successfully trained jointly and can now be queried from either domain!")
// 	// Final analyses
 	fmt.Println("\nRunning full post-training analysis...\n")
//
// 	// Protein folding
 	protAttnFinal = core.ComputeAttention(proteinData, true, nil)
 	AnalyzePhaseSeparation(protAttnFinal, proteinData, proteinLabels) // you'll need to save labels from GenerateDiverseBioBatch
//
// 	// Quantum
 	quantAttnFinal = core.ComputeAttention(quantumData, false, nil)
 	AnalyzeQuantumClustering(quantAttnFinal, quantumOps)

fmt.Println("\n=== PROTEIN SECONDARY STRUCTURE PREDICTION ===")
testProtein := proteinData[0]  // Use first protein from batch
ss := Predict8StateSecondaryStructure(core, testProtein)
fmt.Print("Predicted: ")
for i, s := range ss {
    if i > 0 && i%10 == 0 {
        fmt.Print(" ")
    }
    fmt.Print(s)
}
fmt.Println("\n(H=Helix, E=Sheet, C=Coil)")


}

func CheckCommutation(s1, s2 string) bool {
	ac := 0
	l := min(len(s1), len(s2))
	for i := 0; i < l; i++ {
		if s1[i] != 'I' && s2[i] != 'I' && s1[i] != s2[i] {
			ac++
		}
	}
	return ac%2 == 0
}
