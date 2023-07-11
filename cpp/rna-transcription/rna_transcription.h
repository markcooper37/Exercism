#if !defined(RNA_TRANSCRIPTION_H)
#define RNA_TRANSCRIPTION_H
#include <string>

namespace rna_transcription {
    std::string to_rna(std::string input);
    char to_rna(char input);
}  // namespace rna_transcription

#endif // RNA_TRANSCRIPTION_H