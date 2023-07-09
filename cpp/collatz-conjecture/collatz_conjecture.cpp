#include "collatz_conjecture.h"
#include <stdexcept>

namespace collatz_conjecture {
    int steps(int num) {
        if (num < 1) {
            throw std::domain_error("input must be positive");
        } else if (num == 1) {
            return 0;
        } else {
            return 1 + (num % 2 == 0 ? steps(num / 2) : steps(3 * num + 1));
        }
    }
}  // namespace collatz_conjecture
