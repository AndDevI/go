// Package weather provides a simple weather forecast service.
// It stores the last reported location and weather condition so other
// parts of the program can read what was most recently forecast.
package weather

// CurrentCondition stores the most recently reported weather condition
// (for example: "sunny", "rainy") set by Forecast.
var CurrentCondition string

// CurrentLocation stores the most recently reported city/location
// (for example: "São Paulo") set by Forecast.
var CurrentLocation string

// Forecast updates CurrentLocation and CurrentCondition with the provided
// values and returns a formatted forecast string for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
