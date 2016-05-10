package main

import (
	"github.com/baishuai/golda"
	"fmt"
	"time"
)

func main() {
	c := golda.Load("./data.txt")
	//c := golda.LoadDir("/Users/Bai/Documents/LDA4j/data/mini/")

	fmt.Println("Corpus size", len(c.GetDocuments()))

	g := golda.LdaGibbsSampler(c.GetDocuments(), c.GetVocabularySize())

	for i := 2; i <= 20; i++ {
		start := time.Now()
		g.Gibbs(i, 2.0, 0.5)
		phi := g.GetPhi()
		topics := golda.Translate(phi, c.GetVocabulary(), 10)
		golda.Explain(topics)
		fmt.Printf("i= %d, LogLikelihood: %f", i, g.LogLike())
		cost := time.Now().Sub(start)
		fmt.Printf(" time %f s\n", float64(cost.Nanoseconds()) / (1000.0 * 1000.0 * 1000.0))
	}
}

//i= 2, LogLikelihood: 13996.458905 time 14.910211 s
//i= 3, LogLikelihood: 15746.855800 time 15.264081 s
//i= 4, LogLikelihood: 23942.283241 time 16.229406 s
//i= 5, LogLikelihood: 35255.520423 time 18.232099 s
//i= 6, LogLikelihood: 49038.994690 time 19.854350 s
//i= 7, LogLikelihood: 64431.877574 time 22.401362 s
//i= 8, LogLikelihood: 81269.240974 time 23.371389 s
//i= 9, LogLikelihood: 99299.643388 time 25.838947 s
//i= 10, LogLikelihood: 117674.383103 time 27.944604 s
//i= 11, LogLikelihood: 136989.611163 time 28.513248 s
//i= 12, LogLikelihood: 156557.026968 time 28.224046 s
//i= 13, LogLikelihood: 176585.378980 time 29.519388 s
//i= 14, LogLikelihood: 196957.460613 time 31.627025 s
//i= 15, LogLikelihood: 218004.989120 time 32.440533 s
//i= 16, LogLikelihood: 238452.936949 time 34.286301 s
//i= 17, LogLikelihood: 259745.006693 time 35.202098 s
//i= 18, LogLikelihood: 281477.153226 time 34.665624 s
//i= 19, LogLikelihood: 303445.533222 time 35.699217 s
//i= 20, LogLikelihood: 325564.661640 time 39.295576 s
