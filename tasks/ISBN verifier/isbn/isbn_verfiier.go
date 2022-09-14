package isbn

import (
	"math"
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	//panic("Please implement the IsValidISBN function")
	isbn = strings.ReplaceAll(strings.TrimSpace(isbn), "-", "")

	if len(isbn) != 10 {
		return false
	}

	totals := 0
	for i, d := range strings.Split(isbn, "") {
		if d == "X" && i == len(isbn)-1 {
			totals += 10 * (10 - i)
		} else {
			num, err := strconv.Atoi(d)

			if err != nil {
				return false
			}

			totals += num * (10 - i)
		}
	}

	return math.Mod(float64(totals), 11) == 0.0
}
