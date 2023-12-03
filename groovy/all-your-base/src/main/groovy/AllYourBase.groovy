class AllYourBase {

    def digits
    def inputBase

    AllYourBase(inputBase, digits) {
        if (inputBase <= 1) {
            throw new ArithmeticException("input base must be greater than 1")
        }
        this.digits = digits
        this.inputBase = inputBase
    }

    def rebase(outputBase) {
        if (outputBase <= 1) {
            throw new ArithmeticException("output base must be greater than 1")
        }
        def number = 0
        for (digit in digits) {
            if (digit < 0 || digit >= inputBase) {
                throw new ArithmeticException("invalid digit")
            }
            number = inputBase * number + digit
        }
        def output = []
        while (number > 0) {
            output.add(number % outputBase)
            number = number.intdiv(outputBase)
        }
        if (output == []) {
            output = [0]
        }
        return output.reverse()
    }
}