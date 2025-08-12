// Package weather provides tools to forecast the current weather condition.
package weather

// CurrentCondition represents the current wather condition.
var CurrentCondition string
// CurrentLocation represents current location.
var CurrentLocation string

// Forecast takes takes two strings city and location and returns a string formatting the current weather condition in the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
