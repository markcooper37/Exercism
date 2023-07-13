#include "triangle.h"
#include <stdexcept>

namespace triangle {
    bool valid(double a, double b, double c) {
        return a > 0 && b > 0 && c > 0 && a + b > c && a + c > b && b + c > a;
    }

    flavor kind(double a, double b, double c) {
        if (!valid(a, b, c)) throw std::domain_error("invalid side lengths");
        else if (a == b && b == c) return equilateral;
        else if (a == b || b == c || a == c) return isosceles;
        else return scalene;
    }
}  // namespace triangle
