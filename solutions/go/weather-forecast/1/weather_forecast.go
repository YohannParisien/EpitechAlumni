//Package weather for a specific location.
package weather

var (
    //CurrentCondition : Weather Condition.
	CurrentCondition string
    //CurrentLocation : Location for the forecast.
	CurrentLocation  string
)

// Forecast : Give the forecast for a city, the condition is passed a second parameter.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
