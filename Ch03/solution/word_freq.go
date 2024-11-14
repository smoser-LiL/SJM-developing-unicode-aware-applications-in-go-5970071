package freq

import (
	"regexp"
)

var wordRe = regexp.MustCompile(`\pL+`)

// WordFreq returns a map of word to how many times it appears in the text.
// WordFreq("one two one two three") -> {"one": 2, "two": 2, "three": 1}
func WordFreq(text string) map[string]int {
	freqs := make(map[string]int)
	for _, w := range wordRe.FindAllString(text, -1) {
		freqs[w]++
	}

	return freqs
}
