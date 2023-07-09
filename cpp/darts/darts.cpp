#include "darts.h"
#include <cmath>

namespace darts {
    int score(double x, double y) {
        double distance = std::sqrt(std::pow(x, 2) + std::pow(y, 2));
        if (distance <= 1) return 10;
        else if (distance <= 5) return 5;
        else if (distance <= 10) return 1;
        else return 0;
    }
} // namespace darts