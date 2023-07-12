#if !defined(BINARY_SEARCH_H)
#define BINARY_SEARCH_H
#include <vector>

namespace binary_search {
    int find(std::vector<int> array, int value);
    int position(std::vector<int> array, int value);
}  // namespace binary_search

#endif // BINARY_SEARCH_H