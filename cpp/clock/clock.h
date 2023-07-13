#if !defined(CLOCK_H)
#define CLOCK_H
#include <string>

using namespace std;

namespace date_independent {
    class clock {
        int minutes;
    public:
        clock(int minutes);
        static clock at(int hours, int mins) { return clock(hours * 60 + mins); }
        clock plus(int mins) const { return clock(minutes + mins); }  
        operator std::string() const;
        bool operator ==(clock clk) const { return minutes == clk.minutes; }
        bool operator !=(clock clk) const { return minutes != clk.minutes; }
    };
}  // namespace date_independent

#endif // CLOCK_H