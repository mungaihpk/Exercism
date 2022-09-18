package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	//panic("Please implement the Nth function")
	var err error

	if n < 1 {
		return -1, errors.New("Invalid input.")
	}

	count := 0
	prime := 0
	for i := 2; ; i++ {
		if IsPrime(i) {
			count++
			prime = i

			if count == n {
				break
			}
		}
	}

	return prime, err
}

func IsPrime(num int) bool {
	if num == 1 {
		return true
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
