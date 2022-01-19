package sieve

func Sieve(limit int) []int {
    numbers := map[int]bool{}
	for i := 2; i <= limit; i++ {
        numbers[i] = true
    }
	for i := 2; i <= limit; i++ {
        for j := 2; i * j <= limit; j++ {
            numbers[i * j] = false
        }
    }
	primes := []int{}
	for i := 2; i <= limit; i++ {
        if numbers[i] == true {
            primes = append(primes, i)
        }
    }
	return primes
}
