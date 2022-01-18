// Package gigasecond contains a function to determine the 
// moment a gigasecond after a given time
package gigasecond

import "time"

// AddGigasecond takes a given time and determines the moment 
// that would be after a gigasecond has passed
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
