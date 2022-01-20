// Package leap implements a function to check if a given year is a leap year.
package leap

// Function IsLeapYear checks whether a given year is a leap year.
func IsLeapYear(year int) bool {
	if year % 4 == 0 {
        if year % 400 == 0 {
            return true
        } else if year % 100 == 0 {
        	 return false
        } else {
        	return true
        }
    }
	return false
}
