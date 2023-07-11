#include "perfect_numbers.h"
#include <stdexcept>

namespace perfect_numbers {
    Classification classify(int num) {
        if (num < 1) throw std::domain_error("input must be positive");
        int aliquot_sum{0};
        for (int i = 1; i <= num / 2; i++) {
            if (num % i == 0) aliquot_sum += i;
        }
        if (aliquot_sum == num) return perfect;
        else if (aliquot_sum < num) return deficient;
        else return abundant;
    }
}  // namespace perfect_numbers
