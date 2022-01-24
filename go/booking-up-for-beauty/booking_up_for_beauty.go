package booking

import (
    "time"
    "fmt"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    newDate, err := time.Parse("1/02/2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return newDate
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
    newDate, err := time.Parse("January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return newDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
    newDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return newDate.Hour() >= 12 && newDate.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
    newDate, err := time.Parse("1/2/2006 15:04:05", date)
    if err != nil {
        panic(err)
    }
	return fmt.Sprintf("You have an appointment on %s", newDate.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
} 
