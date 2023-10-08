class Darts {
    int score(double xOfDart, double yOfDart) {
        double distanceSquared = xOfDart * xOfDart + yOfDart * yOfDart;
        if (distanceSquared <= 1) {
            return 10;
        } else if (distanceSquared <= 25) {
            return 5;
        } else if (distanceSquared <= 100) {
            return 1;
        } else {
            return 0;
        }
    }
}
