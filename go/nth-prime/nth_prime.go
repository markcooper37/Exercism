package prime

func Nth(n int) (int, bool) {
	if n < 1 {
        return 0, false
    }
	primes := []int{2}
    numberToCheck := 3
	for len(primes) < n {
        isPrime := true
        for _, prime := range primes {
            if numberToCheck % prime == 0 {
                isPrime = false
                numberToCheck++
                break
            }
        }
    	if isPrime {
            primes = append(primes, numberToCheck)
        	numberToCheck++
        }
    }
	return primes[n - 1], true
}
