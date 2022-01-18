package prime

func Factors(n int64) []int64 {
	factors := []int64{}
	for i := int64(2); i <= n; i++ {
        for n % i == 0 {
            factors = append(factors, i)
            n = n / i
        }
    	if n == 1 {
            break
        }
    }
	return factors
}
