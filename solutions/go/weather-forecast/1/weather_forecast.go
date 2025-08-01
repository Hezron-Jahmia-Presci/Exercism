// Package weather provides tools to forecast the weather accross various cities.
package weather

// CurrentCondition represents a certain weather condition.
var CurrentCondition string

//CurrentLocation represents a certain city.
var CurrentLocation string

// Forecast takes city and condition as its arguments and sets their values in CurrentCondition and CurrentLocation and returrns a string showing the current city location and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
