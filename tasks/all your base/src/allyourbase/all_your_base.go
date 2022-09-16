package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	//panic("Please implement the ConvertToBase function")
	var err error
	err = nil
	var new_num []int

	is_all_zeros := true

	if inputBase < 2 {
		err = errors.New("input base must be >= 2")
		return inputDigits, err
	} else if outputBase < 2 {
		err = errors.New("output base must be >= 2")
		return inputDigits, err
	}

	num := 0
	size := len(inputDigits)
	for i, d := range inputDigits {
		if d >= inputBase || d < 0 {
			return new_num, errors.New("all digits must satisfy 0 <= d < input base")
		} else if d > 0 {
			is_all_zeros = false
		}
		num += d * int(math.Pow(float64(inputBase), float64(size-1-i)))
	}

	if is_all_zeros {
		new_num = append(new_num, 0)
		return new_num, errors.New("all digits must satisfy 0 <= d < input base")
	}

	for num > 0 {
		if num < outputBase {
			new_num = append([]int{num}, new_num...)
			num = 0
			continue
		}

		rem := num % outputBase
		num = num / outputBase
		new_num = append([]int{rem}, new_num...)
	}

	if len(new_num) == 0 {
		new_num = append(new_num, 0)
	}

	return new_num, err
}
