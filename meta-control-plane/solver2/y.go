package main

import (
//	"fmt"
	"math"
//	"math/rand"
//	"strings"
)

// ================================================================
// PART 5: RIEMANNIAN NESTEROV ACCELERATED GRADIENT (RNAG)
// ================================================================

type RiemannianNAG struct {
	LR           float32
	Momentum     float32
	TimeStep     int
	
	Params       []*[][]float32
	Gradients    []*[][]float32
	
	MomentumBufs [][][]float32 // Momentum vector (m_t)
	
	Manifold     *StiefelManifold
}

func NewRiemannianNAG(lr, momentum float32) *RiemannianNAG {
	return &RiemannianNAG{
		LR:       lr,
		Momentum: momentum,
		Manifold: &StiefelManifold{},
		TimeStep: 0,
	}
}

func (opt *RiemannianNAG) AddParameter(param, grad *[][]float32) {
	opt.Params = append(opt.Params, param)
	opt.Gradients = append(opt.Gradients, grad)
	
	n, p := len(*param), len((*param)[0])
	
	// Initialize momentum buffer
	buf := make([][]float32, n)
	for i := range buf {
		buf[i] = make([]float32, p)
	}
	opt.MomentumBufs = append(opt.MomentumBufs, buf)
}

// Replace Step() in RiemannianNAG in y.go

func (opt *RiemannianNAG) Step() {
	opt.TimeStep++
	
	for i := range opt.Params {
		W_t := *opt.Params[i]
		G_t := *opt.Gradients[i]
		n, p := len(W_t), len(W_t[0])
		m_t := opt.MomentumBufs[i]
		
		if n >= p {
			// 1. LOOK-AHEAD with proper transport
			// W_temp = Retraction(W_t, -Momentum * m_t)
			negMomentum := make([][]float32, n)
			for j := range negMomentum {
				negMomentum[j] = make([]float32, p)
				for k := range negMomentum[j] {
					negMomentum[j][k] = -opt.Momentum * m_t[j][k]
				}
			}
			W_temp := opt.Manifold.RetractionCayley(W_t, negMomentum, 1.0)
			
			// 2. GRADIENT at look-ahead point (should be computed externally)
			// Using G_t as placeholder
			G_temp_Manifold := opt.Manifold.ProjectTangent(W_temp, G_t)
			
			// 3. TRANSPORT momentum to W_temp
			m_trans := opt.Manifold.ParallelTransport(W_t, W_temp, m_t)
			
			// 4. UPDATE MOMENTUM with transported m_t
			m_new := make([][]float32, n)
			for j := range m_new {
				m_new[j] = make([]float32, p)
				for k := range m_new[j] {
					m_new[j][k] = G_temp_Manifold[j][k] + opt.Momentum*m_trans[j][k]
				}
			}
			
			// 5. STEP with proper retraction
			negStep := make([][]float32, n)
			for j := range negStep {
				negStep[j] = make([]float32, p)
				for k := range negStep[j] {
					negStep[j][k] = -opt.LR * m_new[j][k]
				}
			}
			
			// Use Cayley for better performance
			W_next := opt.Manifold.RetractionCayley(W_t, negStep, 1.0)
			
			// 6. Update
			for j := range W_next {
				copy(W_t[j], W_next[j])
				copy(m_t[j], m_new[j])
			}
		}
	}
}

// ================================================================
// PART 4: RIEMANNIAN ADAM OPTIMIZER (Replaces RiemannianSGD)
// ================================================================

type RiemannianADAM struct {
	LR           float32 // Learning Rate
	Beta1        float32 // Decay rate for first moment estimate (m)
	Beta2        float32 // Decay rate for second moment estimate (v)
	Epsilon      float32 // Epsilon for stability
	TimeStep     int     // Global time step counter
	
	Params       []*[][]float32
	Gradients    []*[][]float32
	
	// Moment buffers (same shape as Params)
	MomentM      [][][]float32 // First moment (mean) buffer
	MomentV      [][][]float32 // Second moment (variance) buffer
	
	Manifold     *StiefelManifold
}

func NewRiemannianADAM(lr, beta1, beta2, epsilon float32) *RiemannianADAM {
	return &RiemannianADAM{
		LR:       lr,
		Beta1:    beta1,
		Beta2:    beta2,
		Epsilon:  epsilon,
		Manifold: &StiefelManifold{},
		TimeStep: 0,
	}
}

func (opt *RiemannianADAM) AddParameter(param, grad *[][]float32) {
	opt.Params = append(opt.Params, param)
	opt.Gradients = append(opt.Gradients, grad)
	
	n, p := len(*param), len((*param)[0])
	
	// Initialize moment buffers to zero
	m := make([][]float32, n)
	v := make([][]float32, n)
	for i := range m {
		m[i] = make([]float32, p)
		v[i] = make([]float32, p)
	}
	opt.MomentM = append(opt.MomentM, m)
	opt.MomentV = append(opt.MomentV, v)
}

func (opt *RiemannianADAM) Step() {
	opt.TimeStep++
	t := float64(opt.TimeStep)
	
	// Bias correction factors
	m_correction := float32(1.0 / (1.0 - math.Pow(float64(opt.Beta1), t)))
	v_correction := float32(1.0 / (1.0 - math.Pow(float64(opt.Beta2), t)))

	for i := range opt.Params {
		param := *opt.Params[i]
		grad := *opt.Gradients[i]
		n, p := len(param), len(param[0])
		
		if n >= p {
			// 1. PROJECT: Get the Riemannian gradient (tangent vector)
			gradManifold := opt.Manifold.ProjectTangent(param, grad)
			
			m := opt.MomentM[i]
			v := opt.MomentV[i]
			
			// 2. UPDATE MOMENTS (in tangent space)
			stepDirection := make([][]float32, n)
			for j := range param {
				stepDirection[j] = make([]float32, p)
				for k := range param[j] {
					g_man := gradManifold[j][k]
					
					// Update m: m_t = beta1 * m_{t-1} + (1-beta1) * g_manifold
					m[j][k] = opt.Beta1*m[j][k] + (1.0-opt.Beta1)*g_man
					
					// Update v: v_t = beta2 * v_{t-1} + (1-beta2) * g_manifold^2
					v[j][k] = opt.Beta2*v[j][k] + (1.0-opt.Beta2)*g_man*g_man
					
					// Compute step direction (ADAM formula)
					m_hat := m[j][k] * m_correction
					v_hat := v[j][k] * v_correction
					
					// xi = m_hat / (sqrt(v_hat) + epsilon)
					stepDirection[j][k] = m_hat / (float32(math.Sqrt(float64(v_hat))) + opt.Epsilon)
				}
			}
			
			// 3. RETRACTION: Move along the step direction and back onto the manifold
			// The retraction moves in the direction of -stepDirection
			negStep := make([][]float32, n)
			for j := range negStep {
				negStep[j] = make([]float32, p)
				for k := range negStep[j] {
					negStep[j][k] = -stepDirection[j][k]
				}
			}
			
			pNew := opt.Manifold.RetractionQR(param, negStep, opt.LR)
			
			// Update the actual parameter
			for j := range param {
				copy(param[j], pNew[j])
			}
		} else {
			// Fallback to Euclidean ADAM for non-Stiefel parameters (not implemented here)
			// ...
		}
	}
}
