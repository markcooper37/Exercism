class SumOfMultiples {

    static int sum(List<Integer> factors, int limit) {
        def multiples = []
        for (factor in factors) {
            if (factor == 0) {
                continue
            }
            for (def i = factor; i < limit; i += factor) {
                multiples.add(i)
            }
        }
        return multiples.toSet().sum() ?: 0
    }
}