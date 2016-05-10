package golda

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Corpus struct {
	vocabulary *Vocabulary
	documents  [][]int
}

func (c *Corpus) GetVocabularySize() int {
	return c.vocabulary.size()
}

func (c *Corpus) GetVocabulary() *Vocabulary {
	return c.vocabulary
}

func (c *Corpus) addDocument(document []string) {

	doc := make([]int, len(document))
	for i, word := range document {
		doc[i] = c.vocabulary.getIdOrC(word, true)
	}
	c.documents = append(c.documents, doc)
}

func (c *Corpus) GetDocuments() [][]int {
	return c.documents
}

func Load(path string) *Corpus {
	c := new(Corpus)
	c.vocabulary = NewVocabulary()
	c.documents = make([][]int, 0)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer file.Close()

	bio := bufio.NewReader(file)
	line, _, err := bio.ReadLine()

	n, err := strconv.Atoi(string(bytes.TrimSpace(line)))
	fmt.Println("doc size", n)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	c.documents = make([][]int, 0, n)

	for {
		line, _, err = bio.ReadLine()
		if err != nil {
			break
		}
		sline := string(line)
		sline = strings.TrimSpace(sline)
		words := strings.Split(sline, " ")
		c.addDocument(words)
	}

	return c
}

func LoadDir(dirPath string) *Corpus {
	c := new(Corpus)
	c.vocabulary = NewVocabulary()
	c.documents = make([][]int, 0)

	//files, err := ioutil.ReadDir(dirPath)
	dir, err := os.Open(dirPath)
	files, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	for _, f := range files {
		file, err := os.Open(dirPath+f)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		defer file.Close()
		bio := bufio.NewReader(file)

		wordList := make([]string, 0)
		for {
			line, _, err := bio.ReadLine()
			if err != nil {
				break
			}
			sline := string(line)
			sline = strings.TrimSpace(sline)
			words := strings.Split(sline, " ")
			wordList = append(wordList, words...)
		}
		c.addDocument(wordList)
	}

	return c
}
