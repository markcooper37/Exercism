package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule string

const (
    First WeekSchedule = "first"
    Second WeekSchedule = "second"
    Third WeekSchedule = "third"
    Fourth WeekSchedule = "fourth"
    Last WeekSchedule = "last"
    Teenth WeekSchedule = "teenth"
)

func DayOfMonth(start int, weekday time.Weekday, month time.Month, year int) int {
    for i := start; i <= start + 6; i++ {
    	if time.Date(year, month, i, 0, 0, 0, 0, time.UTC).Weekday() == weekday {
      		return i
    	}
  	}
	return 0
}

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	switch week {
        case First:
    		return DayOfMonth(1, weekday, month, year)
        case Second:
			return DayOfMonth(8, weekday, month, year)
        case Third:
			return DayOfMonth(15, weekday, month, year)
        case Fourth:
			return DayOfMonth(22, weekday, month, year)
        case Last:
			switch month {
                case time.Month(2):
					if year % 4 == 0 {
             			return DayOfMonth(23, weekday, month, year)
            		} else {
            			return DayOfMonth(22, weekday, month, year)
          			}
                case time.Month(1), time.Month(3), time.Month(5), time.Month(7), time.Month(8), time.Month(10), time.Month(12):
					return DayOfMonth(25, weekday, month, year)
                default:
            		return DayOfMonth(24, weekday, month, year)
            	}
        case Teenth:
    		return DayOfMonth(13, weekday, month, year)
    }
	return 0
}
