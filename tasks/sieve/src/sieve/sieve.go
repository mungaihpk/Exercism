package sieve

func Contains(non_primes []int, non_prime int) bool {
	for _, num := range non_primes {
		if num == non_prime {
			return true
		}
	}
	return false
}

func Doubles(num int, size int) []int {
	doubles := []int{num}
	prod := 0
	count := 2
	for prod < size {
		prod = count * num
		count++
		doubles = append(doubles, prod)
	}

	return doubles
}

func Sieve(limit int) []int {
	//panic("Please implement the Sieve function")
	primes := []int{}
	non_primes := []int{}

	next_prime := 1
	for i := 2; i <= limit; i++ {
		next_prime = next_prime + 1
		if !Contains(non_primes, next_prime) {
			primes = append(primes, next_prime)
			non_primes = append(non_primes, Doubles(next_prime, limit)...)
		}
	}

	return primes
}
