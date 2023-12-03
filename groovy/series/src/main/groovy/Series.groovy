class Series {

    static slices(series, sliceLength) {
        if (sliceLength <= 0) {
            throw new ArithmeticException("slice length must be positive")
        } else if (sliceLength > series.length()) {
            throw new ArithmeticException("slice length cannot be longer than series length")
        }
        def substrings = []
        for (i in 0..series.length()-sliceLength) {
            substrings.add(series.substring(i, i + sliceLength))
        }
        return substrings
    }
}