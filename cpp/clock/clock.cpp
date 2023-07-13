#include "clock.h"
#include <sstream>
#include <iomanip>

namespace date_independent {
    clock::clock(int mins) { minutes = (mins % 1440 + 1440) % 1440; }

    clock::operator string() const {
        ostringstream strstream;
        strstream << setfill('0') << setw(2) << minutes / 60 << ":" << setw(2) << minutes % 60;
        return strstream.str();
    }
}  // namespace date_independent
