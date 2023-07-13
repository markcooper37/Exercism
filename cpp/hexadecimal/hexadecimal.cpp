#include "hexadecimal.h"
#include <string>

namespace hexadecimal {
    int convert(std::string input) {
        int total{0};
        for (int i = 0; i < int(input.length()); i++) {
            total *= 16;
            if (input[i] >= '0' && input[i] <= '9') total += input[i] - '0';
            else if (input[i] >= 'a' && input[i] <= 'f') total += input[i] + 10 - 'a';
            else return 0;
        }
        return total;
    }
}  // namespace hexadecimal
