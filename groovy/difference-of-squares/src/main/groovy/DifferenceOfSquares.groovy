class DifferenceOfSquares {

    int num

    DifferenceOfSquares(num) {
        this.num = num
    }

    def squareOfSum() {
        def sum = 0
        1.upto(this.num) {
            sum += it
        }
        return sum * sum
    }

    def sumOfSquares() {
        def sum = 0
        1.upto(this.num) {
            sum += it * it
        }
        return sum
    }

    def difference() {
        return squareOfSum() - sumOfSquares()
    }

}
