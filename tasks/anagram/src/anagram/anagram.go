package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	//panic("Please implement the Detect function")
	var anagrams []string

	subject_sub := strings.Split(strings.TrimSpace(subject), "")

outer:
	for _, candidate := range candidates {
		lwr_candidate := strings.ToLower(candidate)
		lwr_subject := strings.ToLower(subject)

		if lwr_subject == lwr_candidate {
			continue
		} else if len(candidate) != len(subject) {
			continue
		}

		for _, letter := range subject_sub {
			lwr_letter := strings.ToLower(letter)

			if !strings.Contains(lwr_candidate, lwr_letter) {
				continue outer
			} else if strings.Count(lwr_subject, lwr_letter) > strings.Count(lwr_candidate, lwr_letter) {
				continue outer
			}
		}

		anagrams = append(anagrams, candidate)
	}

	return anagrams
}
