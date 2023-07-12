#include "bob.h"
#include <string>
#include <algorithm>

namespace bob {
    std::string trim(std::string str) {
        size_t start = str.find_first_not_of(" \n\r\t\f\v");
        str = (start == std::string::npos) ? "" : str.substr(start);
        size_t end = str.find_last_not_of(" \n\r\t\f\v");
        return (end == std::string::npos) ? "" : str.substr(0, end + 1);
    }

    std::string upper(std::string str) {
        for (auto & c: str) c = toupper(c);
        return str;
    }

    bool contains_letter(std::string str) {
        return str.find_first_of("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") != std::string::npos;
    }

    std::string hey(std::string str) {
        std::string trimmed = trim(str);
        if (trimmed == "") return "Fine. Be that way!";
        if (contains_letter(trimmed)) {
            if (trimmed[int(trimmed.size())-1] == '?' && trimmed == upper(trimmed)) return "Calm down, I know what I'm doing!";
            if (trimmed == upper(trimmed)) return "Whoa, chill out!";
        }
        if (trimmed[int(trimmed.size())-1] == '?') return "Sure.";
        return "Whatever.";
    }
}  // namespace bob
