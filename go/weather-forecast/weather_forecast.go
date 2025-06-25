// Package weather provides tool to predict current weather condidion.
package weather

// CurrentCondition declaration.
var CurrentCondition string

// CurrentLocation declaration.
var CurrentLocation string

// Forecast takes 2 parametrs and returns current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
