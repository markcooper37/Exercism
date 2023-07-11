#include "secret_handshake.h"
#include <string>
#include <vector>
#include <algorithm>

namespace secret_handshake {
    std::vector<std::string> commands(int num) {
        std::vector<std::string> output {};
        if (num & 1) output.emplace_back("wink");
        if (num & 2) output.emplace_back("double blink");
        if (num & 4) output.emplace_back("close your eyes");
        if (num & 8) output.emplace_back("jump");
        if (num & 16) std::reverse(output.begin(), output.end());
        return output;
    }
}  // namespace secret_handshake
