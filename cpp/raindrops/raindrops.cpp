#include "raindrops.h"

namespace raindrops {
    std::string convert(int num) {
        std::string sound = "";
        if (num % 3 == 0) sound += "Pling";
        if (num % 5 == 0) sound += "Plang";
        if (num % 7 == 0) sound += "Plong";
        if (sound == "") sound = std::to_string(num);
        return sound;
    }
}  // namespace raindrops
