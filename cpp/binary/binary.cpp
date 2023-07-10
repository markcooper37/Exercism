#include "binary.h"
#include <string>

namespace binary {
    int convert(std::string input) {
        int total{0};
        for (int i = 0; i < int(input.length()); i++) {
            total *= 2;
            if (input[i] == '1') total += 1;
            else if (input[i] != '0') return 0;
        }
        return total;
    }
}  // namespace binary
