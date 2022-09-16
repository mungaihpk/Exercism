package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	//panic("Please implement the Transform function")
	new_scores := make(map[string]int)

	for score, letters := range in {
		for _, letter := range letters {
			small := strings.ToLower(letter)
			new_scores[small] = score
		}
	}

	return new_scores
}
