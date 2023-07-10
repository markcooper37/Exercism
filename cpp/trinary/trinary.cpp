#include "trinary.h"
#include <string>

namespace trinary {
    int to_decimal(std::string input) {
        int total{0};
        for (int i = 0; i < int(input.length()); i++) {
            total *= 3;
            if (input[i] == '1') total += 1;
            else if (input[i] == '2') total += 2;
            else if (input[i] != '0') return 0;
        }
        return total;
    }
}  // namespace trinary
