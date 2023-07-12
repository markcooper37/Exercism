#include "prime_factors.h"
#include <vector>

namespace prime_factors {
    std::vector<int> of(int num) {
        std::vector<int> factors {};
        int factor{2};
        while (num > 1) {
            if (num % factor == 0) {
                factors.emplace_back(factor);
                num /= factor;
            } else factor++;
        }
        return factors;
    }
}  // namespace prime_factors
