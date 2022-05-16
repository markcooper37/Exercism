package sieve

func Sieve(limit int) []int {
    notPrime := map[int]bool{}
	for i := 2; i <= limit; i++ {
        for j := 2; i * j <= limit; j++ {
            notPrime[i * j] = true
        }
    }
	primes := []int{}
	for i := 2; i <= limit; i++ {
        if _, ok := notPrime[i]; !ok {
            primes = append(primes, i)
        }
    }
	return primes
}
