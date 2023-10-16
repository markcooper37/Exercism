import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {
    private LocalDateTime dateTime;
    
    public Gigasecond(LocalDate moment) {
        this.dateTime = moment.atStartOfDay().plusSeconds(1000000000);
    }

    public Gigasecond(LocalDateTime moment) {
        this.dateTime = moment.plusSeconds(1000000000);
    }

    public LocalDateTime getDateTime() {
        return this.dateTime;
    }
}
