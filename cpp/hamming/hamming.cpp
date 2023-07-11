#include "hamming.h"
#include <string>
#include <stdexcept>

namespace hamming {
    int compute(std::string strand_1, std::string strand_2) {
        if (strand_1.size() != strand_2.size()) {
            throw std::domain_error("strands must have same length");
        }
        int count{0};
        for (int i = 0; i < int(strand_1.size()); i++) {
            if (strand_1[i] != strand_2[i]) count++;
        }
        return count;
    }
}  // namespace hamming
