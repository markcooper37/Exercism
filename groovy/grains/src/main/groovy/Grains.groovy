class Grains {
    static square(num) {
        if (num < 1) {
            throw new ArithmeticException("input cannot be negative")
        } else if (num > 64) {
            throw new ArithmeticException("input cannot be greater than 64")
        }
        return new BigInteger("0").setBit(num - 1)
    }

    static total() {
        def sum = new BigInteger("0")
        1.upto(64) {sum = sum.add(square(it))}
        return sum
    }
}
