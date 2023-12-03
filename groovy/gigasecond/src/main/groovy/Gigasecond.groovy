import java.time.LocalDate;
import java.time.LocalDateTime;

class Gigasecond {

    static add(LocalDate moment) {
        return add(moment.atStartOfDay())
    }

    static add(LocalDateTime moment) {
        return moment + 1000000000
    }

}
