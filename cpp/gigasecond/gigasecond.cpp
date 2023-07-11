#include "gigasecond.h"
#include "boost/date_time/posix_time/posix_time.hpp"

using namespace boost::posix_time;
namespace gigasecond {
    ptime advance(ptime input) {
        return input + seconds(1000000000);
    }
}  // namespace gigasecond
