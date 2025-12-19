// analysis.go - Full analysis and visualization for the unified Omega model
// Add this to your joint training program (or import as separate file)

package main

import (
	"fmt"
	"math"
	"strings"
)

// ====================== PROTEIN FOLDING ANALYSIS ======================

func AnalyzePhaseSeparation(attn [][][][]float32, proteinData [][][]float32, labels []string) {
	heads := len(attn[0])
	seqLen := len(attn[0][0][0]) // assumes all batches have same seq len

	fmt.Printf("\n=== PROTEIN FOLDING HEAD SPECIALIZATION ANALYSIS ===\n")
	fmt.Printf("%-4s | %-12s | %-15s | %-15s | %s\n", "Head", "Helix Score", "Sheet Score", "Mixed Score", "Role")
	fmt.Println(strings.Repeat("-", 68))

	// Identify representative batches
	alphaBatch := -1
	betaBatch := -1
	mixedBatch := -1
	for b, label := range labels {
		if label == "ALPHA_BUNDLE" && alphaBatch == -1 {
			alphaBatch = b
		}
		if label == "BETA_BARREL" && betaBatch == -1 {
			betaBatch = b
		}
		if label == "MIXED" && mixedBatch == -1 {
			mixedBatch = b
		}
	}

	countHelix := 0
	countSheet := 0
	countMixed := 0

	for h := 0; h < heads; h++ {
		helixScore := computeLocalCorrelationScore(attn[alphaBatch][h], seqLen)
		sheetScore := computeAntiDiagonalScore(attn[betaBatch][h], seqLen)
		mixedScore := computeMixedScore(attn[mixedBatch][h], seqLen)

		role := "Unspecialized"
		if helixScore > 0.28 && helixScore > sheetScore*1.3 {
			role = "HELIX SPECIALIST 🧬"
			countHelix++
		} else if sheetScore > 0.28 && sheetScore > helixScore*1.3 {
			role = "SHEET SPECIALIST 📄"
			countSheet++
		} else if mixedScore > 0.35 {
			role = "MIXED REGIME 🌪️"
			countMixed++
		}

		fmt.Printf(" #%d  |   %.3f      |   %.3f        |   %.3f        | %s\n",
			h, helixScore, sheetScore, mixedScore, role)

		// Print heatmap for the first discovered specialist of each type
		if role == "HELIX SPECIALIST 🧬" && countHelix == 1 {
			fmt.Printf("\n[Head %d] Alpha-Helix Pattern (local i → i+4)\n", h)
			PrintAttentionHeatmap(attn[alphaBatch][h], "HELIX")
		}
		if role == "SHEET SPECIALIST 📄" && countSheet == 1 {
			fmt.Printf("\n[Head %d] Beta-Sheet Pattern (anti-diagonal pairing)\n", h)
			PrintAttentionHeatmap(attn[betaBatch][h], "SHEET")
		}
	}

	fmt.Printf("\nSummary: %d Helix heads | %d Sheet heads | %d Mixed/Unspecialized\n", countHelix, countSheet, heads-countHelix-countSheet)
	if countHelix > 0 && countSheet > 0 {
		fmt.Println("SUCCESS: Clear phase separation achieved! Model understands distinct physical regimes.")
	} else {
		fmt.Println("NOTE: Further training or higher capacity may be needed for full bifurcation.")
	}
}

func computeLocalCorrelationScore(attn [][]float32, seqLen int) float32 {
	score := float32(0.0)
	total := float32(0.0)
	for i := 0; i < seqLen; i++ {
		for j := 0; j < seqLen; j++ {
			p := attn[i][j]
			total += p
			dist := math.Abs(float64(i - j))
			if dist >= 3 && dist <= 5 { // classic helix: i to i+4
				score += p
			}
		}
	}
	if total > 0 {
		return score / total
	}
	return 0
}

func computeAntiDiagonalScore(attn [][]float32, seqLen int) float32 {
	score := float32(0.0)
	total := float32(0.0)
	for i := 0; i < seqLen; i++ {
		for j := 0; j < seqLen; j++ {
			p := attn[i][j]
			total += p
			antiDist := math.Abs(float64((seqLen - 1 - j) - i))
			dist := math.Abs(float64(i - j))
			if antiDist < 5 && dist > float64(seqLen/4) { // strong distant anti-diagonal
				score += p
			}
		}
	}
	if total > 0 {
		return score / total
	}
	return 0
}

func computeMixedScore(attn [][]float32, seqLen int) float32 {
	score := float32(0.0)
	total := float32(0.0)
	for i := 0; i < seqLen; i++ {
		for j := 0; j < seqLen; j++ {
			p := attn[i][j]
			total += p
			dist := math.Abs(float64(i - j))
			if (dist >= 3 && dist <= 5) || (dist > float64(seqLen/4)) {
				score += p
			}
		}
	}
	if total > 0 {
		return score / total
	}
	return 0
}

func PrintAttentionHeatmap(attn [][]float32, title string) {
	size := len(attn)
	scale := 2 // downsample for readability
	fmt.Printf("\n--- %s Attention Heatmap (↓ query position, → key position) ---\n", title)
	fmt.Print("   ")
	for j := 0; j < size; j += 8 {
		fmt.Printf("%8d", j)
	}
	fmt.Println()

	for i := 0; i < size; i += scale {
		fmt.Printf("%2d ", i)
		for j := 0; j < size; j += scale {
			avg := float32(0.0)
			count := 0
			for di := 0; di < scale && i+di < size; di++ {
				for dj := 0; dj < scale && j+dj < size; dj++ {
					avg += attn[i+di][j+dj]
					count++
				}
			}
			val := avg / float32(count)
			char := "·"
			if val > 0.15 {
				char = "."
			}
			if val > 0.30 {
				char = "o"
			}
			if val > 0.50 {
				char = "O"
			}
			if val > 0.70 {
				char = "█"
			}
			if val > 0.90 {
				char = "▓"
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
	fmt.Println()
}

// ====================== QUANTUM CLUSTERING ANALYSIS ======================

func AnalyzeQuantumClustering(attn [][][][]float32, ops [][]PauliOperator) {
	numBatches := len(attn)
	heads := len(attn[0])

	fmt.Printf("\n=== QUANTUM OPERATOR CLUSTERING ANALYSIS ===\n")

	for b := 0; b < numBatches; b++ {
		fmt.Printf("\nHamiltonian Batch %d (%d operators):\n", b+1, len(ops[b]))
		fmt.Printf("%-4s | %-12s | %-10s | %-10s | %s\n", "Head", "Avg Entropy", "Purity", "Energy Focus", "Interpretation")
		fmt.Println(strings.Repeat("-", 60))

		for h := 0; h < heads; h++ {
			headAttn := attn[b][h]
			numOps := len(headAttn)

			// Entropy
			ent := float32(0.0)
			for i := 0; i < numOps; i++ {
				rowEnt := float32(0.0)
				for j := 0; j < numOps; j++ {
					p := headAttn[i][j]
					if p > 1e-9 {
						rowEnt -= p * float32(math.Log(float64(p)))
					}
				}
				ent += rowEnt
			}
			ent /= float32(numOps)

			// Commuting Purity
			purity := float32(0.0)
			count := 0
			for i := 0; i < numOps; i++ {
				maxJ := -1
				maxP := float32(0.0)
				for j := 0; j < numOps; j++ {
					if headAttn[i][j] > maxP && i != j {
						maxP = headAttn[i][j]
						maxJ = j
					}
				}
				if maxJ != -1 {
					if CheckCommutation(ops[b][i].PauliString, ops[b][maxJ].PauliString) {
						purity += maxP
					} else {
						purity -= maxP * 0.5 // mild penalty
					}
					count++
				}
			}
			if count > 0 {
				purity /= float32(count)
			}

			// Energy Concentration
			energyFocus := float32(0.0)
			for j := 0; j < numOps; j++ {
				incoming := float32(0.0)
				for i := 0; i < numOps; i++ {
					incoming += headAttn[i][j]
				}
				energyFocus += incoming * float32(math.Abs(float64(ops[b][j].Coefficient)))
			}

			interpretation := "Diffuse"
			if ent < 0.7 && purity > 0.6 {
				interpretation = "STRONG COMMUTING CLUSTER"
			} else if ent < 0.9 && energyFocus > 1.5 {
				interpretation = "ENERGY-FOCUSED"
			} else if ent < 1.0 {
				interpretation = "SHARPENED"
			}

			fmt.Printf(" #%d  |   %.3f      |   %.3f    |   %.3f    | %s\n",
				h, ent, purity, energyFocus, interpretation)

			// Print one example sharp head
			if interpretation == "STRONG COMMUTING CLUSTER" && b == 0 {
				fmt.Printf("\n[Head %d] Sharp Commuting Cluster Attention Matrix:\n", h)
				printOperatorAttentionMatrix(headAttn, ops[b])
			}
		}
	}
}

func printOperatorAttentionMatrix(attn [][]float32, ops []PauliOperator) {
	n := len(attn)
	fmt.Print("     ")
	for j := 0; j < n; j++ {
		fmt.Printf("%8s ", ops[j].PauliString)
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%4s ", ops[i].PauliString)
		for j := 0; j < n; j++ {
			val := attn[i][j]
			char := "·"
			if val > 0.2 {
				char = "."
			}
			if val > 0.5 {
				char = "o"
			}
			if val > 0.8 {
				char = "█"
			}
			fmt.Printf("%8s", char)
		}
		fmt.Printf("  (coeff=%.2f)\n", ops[i].Coefficient)
	}
	fmt.Println()
}

// ====================== INTEGRATION INTO MAIN ======================

// Add these calls at the end of your main() after joint training:

// 	// Final analyses
// 	fmt.Println("\nRunning full post-training analysis...\n")
//
// 	// Protein folding
// 	protAttnFinal := core.ComputeAttention(proteinData, true, nil)
// 	AnalyzePhaseSeparation(protAttnFinal, proteinData, proteinLabels) // you'll need to save labels from GenerateDiverseBioBatch
//
// 	// Quantum
// 	quantAttnFinal := core.ComputeAttention(quantumData, false, nil)
// 	AnalyzeQuantumClustering(quantAttnFinal, quantumOps)
