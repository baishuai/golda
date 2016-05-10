package golda

type Vocabulary struct {
	word2id map[string]int
	id2word []string
}

func NewVocabulary() *Vocabulary {
	v := new(Vocabulary)
	v.id2word = make([]string, 0, 1024)
	v.word2id = make(map[string]int)
	return v
}

func (v *Vocabulary) getIdOrC(word string, create bool) int {

	id, ok := v.word2id[word]

	if !ok && create {
		id = len(v.id2word)
		v.word2id[word] = id
		v.id2word = append(v.id2word, word)
	}
	return id
}

func (v *Vocabulary) getId(word string) int {

	return v.getIdOrC(word, false)
}

func (v *Vocabulary) getWord(id int) string {
	return v.id2word[id]
}

func (v *Vocabulary) size() int {
	return len(v.word2id)
}
