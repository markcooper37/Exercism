#include "series.h"
#include <vector>
#include <string>
#include <stdexcept>

namespace series {
    std::vector<std::string> slice(std::string digits, int length) {
        if (length < 1 || int(digits.length()) < length) {
            throw std::domain_error("invalid arguments");
        }
        std::vector<std::string> strings {};
        for (int i = 0; i <= int(digits.length()) - length; i++) {
            strings.emplace_back(digits.substr(i, length));
        }
        return strings;
    }
}  // namespace series
