class Hamming {

    def distance(strand1, strand2) {
        if (strand1.length() != strand2.length()) {
            throw new ArithmeticException("strands must have the same length")
        }
        [strand1.toList(), strand2.toList()].transpose().count{it[0] != it[1]}
    }

}