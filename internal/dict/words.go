package dict

import (
	_ "embed"
)

//go:embed words_alpha.txt
var wordsStatic []byte

// The number of words in the file above
const wordsCount = 370105
