package prime

func Factors(n int64) []int64 {
	//panic("Please implement the Factors function")
	var numbers []int64
	var prime_factor int64
	prime_factor = int64(2)

	for n > 1 {
		if n%prime_factor == 0 {
			numbers = append(numbers, prime_factor)
			n = n / prime_factor
			continue
		}
		prime_factor += 1
	}

	return numbers
}
