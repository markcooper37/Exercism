#include "sum_of_multiples.h"
#include <vector>
#include <map>

namespace sum_of_multiples {
    int to(std::vector<int> nums, int limit) {
        std::map<int, bool> multiples {};
        for (int num : nums) {
            for (int i = num; i < limit; i += num) multiples[i] = true;
        }
        int sum{0};
        for (auto &[key, _] : multiples) sum += key;
        return sum;
    }
}  // namespace sum_of_multiples
