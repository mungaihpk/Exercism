package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

var reg_exp = regexp.MustCompile(`[^A-Za-z0-9]+`)

func GetDimensions(size int) (int, int) {
	sqrt := int(math.Sqrt(float64(size)))
	row, col := sqrt, sqrt
	if sqrt*sqrt < size {
		col++
	}

	if row*col < size {
		row++
	}

	return row, col
}

func Encode(pt string) string {
	//panic("Please implement the Encode function")
	pt = reg_exp.ReplaceAllString(strings.ToLower(pt), "")
	size := len(pt)

	row, col := GetDimensions(size)

	slices := [][]string{}

	msg := ""
	sub_str := []string{}
	col_count := 0

	for i := 0; i <= (col * row); i++ {
		l := ""
		if i >= (size - 1) {
			if i == (size - 1) {
				l = pt[i:]
			} else {
				l = " "
			}
		} else {
			l = pt[i : i+1]
		}

		col_count++

		if col_count == col {
			sub_str = append(sub_str, l)
			slices = append(slices, sub_str)
			sub_str = nil
			col_count = 0
		} else {
			sub_str = append(sub_str, l)
		}
	}

	slices = append(slices, sub_str)

	for i := 0; i < col; i++ {
		cypher := ""
		for j := 0; j < row; j++ {
			cypher += slices[j][i]
		}

		msg += cypher
		if i < (col - 1) {
			msg += " "
		}
	}

	return msg
}
