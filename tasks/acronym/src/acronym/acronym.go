package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	//create a regular expression to identify any character not in a-z, A-Z, 0-9,_ and '
	re := regexp.MustCompile(`[^a-zA-Z0-9-' ]+`)

	//use the regular expression, and strings.ReplaceAll() to replace spaces with hyphen (-)
	s = strings.ReplaceAll(re.ReplaceAllString(s, ""), " ", "-")

	//create a slice of strings by spliting the string using the hyphen
	words := strings.Split(s, "-")
	acronym := ""

	for _, l := range words {
		if len(l) > 1 {
			acronym += strings.ToUpper(l[0:1])
		} else {
			acronym += strings.ToUpper(l)
		}
	}

	return acronym
}
