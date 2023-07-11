#include "nth_prime.h"
#include <vector>
#include <algorithm>

namespace nth_prime {
    int nth(int n) {
        if (n < 1) {
            throw std::domain_error("input must be positive");
        }
        std::vector<int> primes {2};
        int number{3};
        while (int(primes.size()) < n) {
            if (std::all_of(primes.begin(), primes.end(), [&number](int i){return number % i != 0;})) {
                primes.emplace_back(number);
            }
            number++;
        }
        return primes.back();
    }
}  // namespace nth_prime
