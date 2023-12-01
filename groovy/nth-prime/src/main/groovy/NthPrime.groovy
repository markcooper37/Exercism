class NthPrime {

    static nth(int n) {
        if (n <= 0) {
            throw new ArithmeticException("input must be positive")
        }
        def primes = [2]
        def current = 3
        while (primes.size < n) {
            if (primes.every{element -> current % element != 0}) {
                primes.add(current)
            }
            current++
        }
        return primes[n-1]
    }

}