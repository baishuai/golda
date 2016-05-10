package golda

import (
	"fmt"
	"math/rand"
)

var (
	_BURN_IN = 100
	_ITERATIONS = 1000
	_SAMPLE_LAG = 10
)

type GibbsSampler struct {
	v         int
	k         int
	m         int
	alpha     float64
	beta      float64

	documents [][]int

	z         [][]int
	nw        [][]int
	nd        [][]int

	nwsum     []int
	ndsum     []int

	thetasum  [][]float64
	phisum    [][]float64
	numstats  int
}

func LdaGibbsSampler(documents [][]int, v int) *GibbsSampler {
	g := new(GibbsSampler)
	g.documents = documents
	g.v = v
	g.m = len(documents)
	return g
}

func (g *GibbsSampler) Gibbs(k int, alpha float64, beta float64) {
	g.k = k
	g.alpha = alpha
	g.beta = beta

	if _SAMPLE_LAG > 0 {
		g.numstats = 0
		g.thetasum = make([][]float64, g.m)
		for i := 0; i < g.m; i++ {
			g.thetasum[i] = make([]float64, g.k)
		}
		g.phisum = make([][]float64, g.k)
		for i := 0; i < g.k; i++ {
			g.phisum[i] = make([]float64, g.v)
		}
	}
	g.initialState()

	//fmt.Println("Sampling ", _ITERATIONS, " iterations with burn-in of ", _BURN_IN, " (B/S=", _THIN_INTERVAL, ").")

	for i := 0; i < _ITERATIONS; i++ {
		for m := 0; m < g.m; m++ {
			for n := 0; n < len(g.z[m]); n++ {
				topic := g.sampleFullConditional(m, n)
				g.z[m][n] = topic
			}
		}

		if (i > _BURN_IN) && (_SAMPLE_LAG > 0) && (i % _SAMPLE_LAG == 0) {
			g.updateParams()
		}
	}
	fmt.Println()
}

func (g *GibbsSampler) initialState() {

	g.nw = make([][]int, g.v)
	for i := 0; i < g.v; i++ {
		g.nw[i] = make([]int, g.k)
	}

	g.nd = make([][]int, g.m)
	for i := 0; i < g.m; i++ {
		g.nd[i] = make([]int, g.k)
	}

	g.nwsum = make([]int, g.k)
	g.ndsum = make([]int, g.m)

	g.z = make([][]int, g.m)
	for m := 0; m < g.m; m++ {
		_n := len(g.documents[m])
		g.z[m] = make([]int, _n)
		for n := 0; n < _n; n++ {
			topic := int(rand.Float64() * float64(g.k))
			g.z[m][n] = topic
			g.nw[g.documents[m][n]][topic]++
			g.nd[m][topic]++
			g.nwsum[topic]++
		}
		g.ndsum[m] = _n
	}
}

func (g *GibbsSampler) sampleFullConditional(m, n int) int {

	topic := g.z[m][n]
	g.nw[g.documents[m][n]][topic]--
	g.nd[m][topic]--
	g.nwsum[topic]--
	g.ndsum[m]--

	p := make([]float64, g.k)
	for k := 0; k < g.k; k++ {
		p[k] = (float64(g.nw[g.documents[m][n]][k]) + g.beta) / (float64(g.nwsum[k]) + float64(g.v) * g.beta) * (float64(g.nd[m][k]) + g.alpha) / (float64(g.ndsum[m]) + float64(g.k) * g.alpha)
	}

	for k := 1; k < g.k; k++ {
		p[k] += p[k - 1]
	}

	u := rand.Float64() * p[g.k - 1]
	for topic = 0; topic < g.k; topic++ {
		if u < p[topic] {
			break
		}
	}

	g.nw[g.documents[m][n]][topic]++
	g.nd[m][topic]++
	g.nwsum[topic]++
	g.ndsum[m]++

	return topic
}

func (g *GibbsSampler) updateParams() {
	for m := 0; m < g.m; m++ {
		for k := 0; k < g.k; k++ {
			g.thetasum[m][k] += (float64(g.nd[m][k]) + g.alpha) / (float64(g.ndsum[m]) + float64(g.k) * g.alpha)
		}
	}
	for k := 0; k < g.k; k++ {
		for w := 0; w < g.v; w++ {
			g.phisum[k][w] += (float64(g.nw[w][k]) + g.beta) / (float64(g.nwsum[k]) + float64(g.v) * g.beta)
		}
	}
	g.numstats++
}

func (g *GibbsSampler) GetPhi() [][]float64 {
	phi := make([][]float64, g.k)

	if _SAMPLE_LAG > 0 {
		for k := 0; k < g.k; k++ {
			phi[k] = make([]float64, g.v)
			for w := 0; w < g.v; w++ {
				phi[k][w] = g.phisum[k][w] / float64(g.numstats)
			}
		}
	} else {
		for k := 0; k < g.k; k++ {
			phi[k] = make([]float64, g.v)
			for w := 0; w < g.v; w++ {
				phi[k][w] = (float64(g.nw[w][k]) + g.beta) / (float64(g.nwsum[k]) + float64(g.v) * g.beta)
			}
		}
	}
	return phi
}

func (g *GibbsSampler)getTheta() [][]float64 {
	theta := make([][]float64, g.m)
	for i := 0; i < g.m; i++ {
		theta[i] = make([]float64, g.k)
	}

	if _SAMPLE_LAG > 0 {
		for m := 0; m < g.m; m++ {
			for k := 0; k < g.k; k++ {
				theta[m][k] = g.thetasum[m][k] / float64(g.numstats)
			}
		}
	} else {
		for m := 0; m < g.m; m++ {
			for k := 0; k < g.k; k++ {
				theta[m][k] = (float64(g.nd[m][k]) + g.alpha) / (float64(g.ndsum[m]) + float64(g.k) * g.alpha)
			}
		}
	}
	return theta
}