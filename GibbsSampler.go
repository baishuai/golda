package golda

import (
	"math"
	"math/rand"
)

const (
	_ITERATIONS = 1000
)

type GibbsSampler struct {
	v     int
	k     int
	m     int
	alpha float64
	beta  float64

	documents [][]int

	z  [][]int
	nw [][]int
	nd [][]int

	nwsum []int
	ndsum []int
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

	g.initialState()

	for i := 0; i < _ITERATIONS; i++ {
		for m := 0; m < g.m; m++ {
			for n := 0; n < len(g.z[m]); n++ {
				topic := g.sampleFullConditional(m, n)
				g.z[m][n] = topic
			}
		}
	}
}

// 随机初始化状态
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

// 计算文档m中第n个词语的主题的完全条件分布，输出最可能的主题
func (g *GibbsSampler) sampleFullConditional(m, n int) int {

	topic := g.z[m][n]
	g.nw[g.documents[m][n]][topic]--
	g.nd[m][topic]--
	g.nwsum[topic]--
	g.ndsum[m]--

	p := make([]float64, g.k)
	for k := 0; k < g.k; k++ {
		p[k] = (float64(g.nw[g.documents[m][n]][k]) + g.beta) / (float64(g.nwsum[k]) + float64(g.v)*g.beta) * (float64(g.nd[m][k]) + g.alpha) / (float64(g.ndsum[m]) + float64(g.k)*g.alpha)
	}

	for k := 1; k < g.k; k++ {
		p[k] += p[k-1]
	}

	u := rand.Float64() * p[g.k-1]
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

// 获取主题——词语矩阵, K x V
func (g *GibbsSampler) getPhi() [][]float64 {
	phi := make([][]float64, g.k)
	for k := 0; k < g.k; k++ {
		phi[k] = make([]float64, g.v)
		for w := 0; w < g.v; w++ {
			phi[k][w] = (float64(g.nw[w][k]) + g.beta) / (float64(g.nwsum[k]) + float64(g.v)*g.beta)
		}
	}
	return phi
}

// 获取文档——主题矩阵, M x K
func (g *GibbsSampler) getTheta() [][]float64 {
	theta := make([][]float64, g.m)
	for i := 0; i < g.m; i++ {
		theta[i] = make([]float64, g.k)
	}

	for m := 0; m < g.m; m++ {
		for k := 0; k < g.k; k++ {
			theta[m][k] = (float64(g.nd[m][k]) + g.alpha) / (float64(g.ndsum[m]) + float64(g.k)*g.alpha)
		}
	}
	return theta
}

// 计算结果的likelihood
func (g *GibbsSampler) LogLike() float64 {

	theta := g.getTheta()
	phi := g.getPhi()
	lik := 0.0
	for d := 0; d < g.m; d++ {
		lw := len(g.documents[d])
		for w := 0; w < lw; w++ {
			word := g.documents[d][w]
			tmp := 0.0
			for k := 0; k < g.k; k++ {
				tmp += theta[d][k] * phi[k][word]
			}
			lik += math.Log(tmp)
		}
	}
	return lik
}
