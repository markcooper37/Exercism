class CollatzConjecture {

    static int steps(int number) {
        if (number < 1) {
            throw new ArithmeticException("number must be positive")
        } else if (number == 1) {
            return 0
        } else if (number % 2 == 0) {
            return 1 + steps(number.intdiv(2))
        } else {
            return 1 + steps(3 * number + 1)
        }
    }
}