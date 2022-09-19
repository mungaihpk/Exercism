package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	words := strings.Split(strings.Trim(phrase, " "), " ")

	if len(words) == 1 {
		words = strings.Split(phrase, ",")
	}

	frequencies := Frequency{}

	for _, word := range words {

		new_word := strings.ToLower(strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		}))
		current_count := frequencies[new_word]

		if new_word == "" {
			continue
		}

		frequencies[new_word] = current_count + 1
	}

	return frequencies
}
