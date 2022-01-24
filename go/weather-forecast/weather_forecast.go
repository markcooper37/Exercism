// Package weather provides a tool to display the weather condition of a particular city.
package weather

// CurrentCondition represents a certain weather condition.
var CurrentCondition string

// CurrentLocation represents a certain city.
var CurrentLocation string

// Forecast returns a string of a city and its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
