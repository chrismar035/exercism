package sieve

func Sieve(n int) []int {
	known := map[int]bool{}
	var primes = []int{}

	for i := 2; i <= n; i++ {
		if known[i] == false {
			primes = append(primes, i)

			for ii := i; ii <= n; ii++ {
				known[i*ii] = true
			}
		}
	}
	return primes
}
