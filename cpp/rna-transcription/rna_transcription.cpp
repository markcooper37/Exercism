#include "rna_transcription.h"
#include <string>

namespace rna_transcription {
    std::string to_rna(std::string input) {
        std::string output = "";
        for (int i = 0; i < int(input.size()); i++) {
            std::string str(1, to_rna(input[i]));
            output += str;
        }
        return output;
    }
    char to_rna(char input) {
            if (input == 'G') return 'C';
            else if (input == 'C') return 'G';
            else if (input == 'T') return 'A';
            else return 'U';
    }
}  // namespace rna_transcription
