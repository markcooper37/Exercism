#include "sieve.h"
#include <vector>
#include <map>

namespace sieve {
    std::vector<int> primes(int num) {
        std::vector<int> prime_values {};
        std::map<int, bool> not_prime {};
        for (int i = 2; i <= num; i++) {
            if (!(not_prime.count(i))) {
                prime_values.emplace_back(i);
                for (int j = i * 2; j <= num; j += i) {
                    not_prime[j] = true;
                }
            }
        }
        return prime_values;
    }

}  // namespace sieve
