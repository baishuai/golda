package main

import (
	"github.com/baishuai/golda"
	"fmt"
	"time"
	"flag"
)

var (
	k int
	alpha float64
	beta float64
)

func init() {
	flag.IntVar(&k, "k", 3, "topic number")
	flag.Float64Var(&alpha, "a", 1.0, "alpha")
	flag.Float64Var(&beta, "b", 0.1, "beta")
	flag.Parse()
	//flag.PrintDefaults()
}

func main() {
	c := golda.Load("./data.txt")

	g := golda.LdaGibbsSampler(c.GetDocuments(), c.GetVocabularySize())

	start := time.Now()
	g.Gibbs(k, alpha, beta)
	topics := g.Translate(c.GetVocabulary(), 10)
	golda.Explain(topics)
	fmt.Printf("k= %d, LogLikelihood: %f", k, g.LogLike())
	cost := time.Now().Sub(start)
	fmt.Printf(" time %f s\n", float64(cost.Nanoseconds()) / (1000.0 * 1000.0 * 1000.0))
}

// go run example.go -k 8
