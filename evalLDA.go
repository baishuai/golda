package golda

import "math"

////def score():
//res = 0.0
//for doc_id in range(0, len(doc)):
//words = doc[doc_id]
//for word_id in range(0, len(words)):
//word = words[word_id]
//tmp = 0.0
//for k in range(0, num_topic):
//tmp += theta[doc_id][k] * phi[k][word]
//res += math.log(tmp)
//return res

func (g *GibbsSampler)LogLike() float64 {

	theta := g.nd
	phi := g.nw
	lik := 0.0
	for z := 0; z < g.k; z++ {
		lik += logMultiBetaVet(phi[z], g.beta)
		lik -= logMultiBeta(g.beta, g.v)
	}
	for m := 0; m < g.m; m++ {
		lik += logMultiBetaVet(theta[m], g.alpha)
		lik -= logMultiBeta(g.alpha, g.k)
	}
	return lik
}


func logMultiBeta(beta float64, k int) float64 {
	return float64(k) * gammaln(beta) - gammaln(beta * float64(k))
}

func logMultiBetaVet(v []int, pend float64) float64 {
	sumV := pend * float64(len(v))
	sumGV := 0.0
	for i := 0; i < len(v); i++ {
		sumV += float64(v[i])
		sumGV += gammaln(float64(v[i]) + pend)
	}
	return sumGV - gammaln(sumV)
}

func gammaln(v float64) float64 {
	ga, _ := math.Lgamma(v)
	return ga
}