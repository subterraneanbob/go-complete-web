package dict

import (
	"bufio"
	"bytes"
	"strings"
)

type Dict struct {
	words []string
}

func NewDict() *Dict {

	d := new(Dict)
	words := make([]string, 0, wordsCount)

	scanner := bufio.NewScanner(bytes.NewReader(wordsStatic))
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	d.words = words
	return d
}

func (d *Dict) Find(prefix string, limit int) []string {

	results := []string{}

	for i := range d.words {
		if strings.HasPrefix(d.words[i], prefix) {
			results = append(results, d.words[i])
			limit--
		}

		if limit == 0 {
			break
		}
	}

	return results
}
