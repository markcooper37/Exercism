#include "grains.h"
#include <cmath>

namespace grains {

long long unsigned int square(int square_number) {
    return pow(2, square_number - 1);
}

long long unsigned int total() {
    int sum{0};
    for (int i = 0; i < 64; i++) {
        sum += square(i);
    }
    return sum;
}

}  // namespace grains
