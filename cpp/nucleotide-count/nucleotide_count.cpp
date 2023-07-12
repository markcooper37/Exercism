#include "nucleotide_count.h"
#include <map>
#include <stdexcept>

namespace nucleotide_count {
    std::map<char, int> count(std::string strand) {
        std::map<char, int> counts {{'A', 0}, {'C', 0}, {'G', 0}, {'T', 0}};
        for (char c: strand) {
            counts[c]++;
            if (counts.size() > 4) throw std::invalid_argument("invalid nucleotide");
        }
        return counts;
    }
}  // namespace nucleotide_count
