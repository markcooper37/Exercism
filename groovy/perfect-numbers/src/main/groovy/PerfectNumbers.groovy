class PerfectNumbers {

    static Classification classify(int num) {
        if (num < 1) {
            throw new ArithmeticException("number must be positive")
        }
        def factorSum = 0
        for (def factor = 1; factor <= num.intdiv(2); factor++) {
            if (num % factor == 0) {
                factorSum += factor
            }
        }
        if (factorSum < num) {
            return Classification.DEFICIENT
        } else if (factorSum == num) {
            return Classification.PERFECT
        } else {
            return Classification.ABUNDANT
        }
    }

}