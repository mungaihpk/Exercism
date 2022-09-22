package sieve

func Sieve(limit int) []int {
	//panic("Please implement the Sieve function")
	primes := []int{}
	non_primes := map[int]bool{}

	for i := 2; i <= limit; i++ {
		if !non_primes[i] {
			primes = append(primes, i)
			for num := 2; num*i <= limit; num++ {
				non_primes[num*i] = true
			}
		}
	}

	return primes
}
