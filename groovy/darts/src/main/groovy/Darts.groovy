class Darts {

    static int score(x, y) {
        def distanceSquared = x * x + y * y
        if (distanceSquared <= 1) {
            return 10
        } else if (distanceSquared <= 25) {
            return 5
        } else if (distanceSquared<= 100) {
            return 1
        } else {
            return 0
        }
    }

}