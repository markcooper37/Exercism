#include "reverse_string.h"

namespace reverse_string {
    std::string reverse_string(std::string input) {
        std::string output = "";
        for (int i = input.length() - 1; i >= 0 ; i--) {
            output += input.at(i);
        }
        return output;
    }
}  // namespace reverse_string
