class PrimeFactors {

    static factors(value) {
        def factors = []
        def factor = 2
        while (value > 1) {
            if (value % factor == 0) {
                factors.add(factor)
                value = value.intdiv(factor)
            } else {
                factor++
            }
        }
        return factors
    }
}