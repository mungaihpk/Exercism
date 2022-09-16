package lsproduct

import (
	"errors"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	//panic("Please implement the LargestSeriesProduct function")
	var err error

	if len(digits) < span {
		return -1, errors.New("invalid input")
	} else if span == 0 {
		return 1, err
	} else if span < 0 {
		return -1, errors.New("span must not be negative")
	}

	digits_slice := strings.Split(digits, "")

	prod := 0

	for index, _ := range digits_slice {
		new_prod := 1

		if index+span > len(digits_slice) {
			break
		}

		for i := index; i < index+span; i++ {
			val, err := strconv.Atoi(digits_slice[i])

			if err != nil {
				return -1, errors.New("digits input must only contain digits")
			}

			new_prod *= val
		}

		if prod < new_prod {
			prod = new_prod
		}
	}

	return int64(prod), err
}
