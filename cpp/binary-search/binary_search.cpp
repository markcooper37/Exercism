#include "binary_search.h"
#include <vector>
#include <stdexcept>

namespace binary_search {
    int find(std::vector<int> array, int value) {
        int result = position(array, value);
        if (result == -1) throw std::domain_error("could not find value");
        else return result;
    }

    int position(std::vector<int> array, int value) {
        int size = array.size();
        if (size == 0) return -1;
        if (array[size / 2] == value) return size / 2;
        else if (array[size / 2] > value) {
            return position({array.begin(), array.begin() + (size / 2)}, value);
        } else {
             int index = position({array.begin() + (size / 2)+1, array.end()}, value);
             if (index != -1) return size / 2 + 1 + index;
             else return -1;
        }
    }
}  // namespace binary_search
