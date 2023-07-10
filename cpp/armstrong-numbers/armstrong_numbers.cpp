#include "armstrong_numbers.h"
#include <vector>
#include <cmath>

namespace armstrong_numbers {
    bool is_armstrong_number(int num) {
        int temp_num = num;
        std::vector<int> digits {};
        int digit_count{0};
        while (temp_num > 0) {
            digit_count++;
            digits.emplace_back(temp_num % 10);
            temp_num /= 10;
        }
        int total{0};
        for (int i = 0; i < digit_count; i++) {
            total += std::pow(digits[i], digit_count);
        }
        return total == num;
    }
}  // namespace armstrong_numbers
