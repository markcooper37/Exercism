class CollatzCalculator {

    int computeStepCount(int start) {
        if (start < 1) {
            throw new IllegalArgumentException("Only natural numbers are allowed");
        } else if (start == 1) {
            return 0;
        } else if (start % 2 == 0) {
            return computeStepCount(start / 2) + 1;
        } else {
            return computeStepCount(3 * start + 1) + 1;
        }
    }

}
