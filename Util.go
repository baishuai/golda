package golda

import (
	"fmt"
	"sort"
)

type _Topic []pair

type pair struct {
	key   string
	value float64
}

func (p _Topic) Len() int {
	return len(p)
}
func (p _Topic) Less(i, j int) bool {
	return p[i].value > p[j].value
}
func (p _Topic) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func Explain(result []_Topic) {
	for i, topic := range result {
		fmt.Printf("topic %d :\n", i)
		for _, pair := range topic {
			fmt.Printf("%s=%.10f\n", pair.key, pair.value)
		}
		fmt.Println()
	}
}

func Translate(phi [][]float64, vocabulary *Vocabulary, limit int) []_Topic {
	if len(phi[0]) < limit {
		limit = len(phi[0])
	}
	lphi := len(phi)
	result := make([]_Topic, 0, lphi)

	for k := 0; k < lphi; k++ {
		rankPair := make(_Topic, 0)
		lphik := len(phi[k])
		for i := 0; i < lphik; i++ {
			rankPair = append(rankPair, pair{key:vocabulary.getWord(i), value:phi[k][i]})
		}
		sort.Sort(rankPair)
		rankPair = rankPair[:limit]
		result = append(result, rankPair)
	}
	return result
}
