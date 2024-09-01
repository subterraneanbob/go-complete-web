package dict

import (
	"bufio"
	"bytes"
)

type MapDict struct {
	words map[string][]string
}

func NewMapDict() *MapDict {

	m := new(MapDict)
	m.words = make(map[string][]string)

	scanner := bufio.NewScanner(bytes.NewReader(wordsStatic))
	for scanner.Scan() {
		word := scanner.Text()
		// Create all possible prefixes starting with the first character of the word,
		// extending it by one character at a time.
		for i := 0; i < len(word); i++ {
			prefix := word[0 : i+1]
			m.words[prefix] = append(m.words[prefix], word)
		}
	}

	return m
}

func (m *MapDict) Find(prefix string, limit int) []string {

	results := m.words[prefix]

	if len(results) > limit {
		return results[:limit]
	}

	return results
}
