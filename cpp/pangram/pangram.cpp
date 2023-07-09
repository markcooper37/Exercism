#include "pangram.h"
#include <string>
#include <algorithm>

namespace pangram {
    bool is_pangram(std::string sentence) {
        std::string letters = "abcdefghijklmnopqrstuvwxyz";
        return std::all_of(letters.begin(), letters.end(), [&sentence](char c){return sentence.find(c)!=std::string::npos || sentence.find(std::toupper(c))!=std::string::npos;});
    }
}  // namespace pangram
