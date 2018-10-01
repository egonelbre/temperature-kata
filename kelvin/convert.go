package temp

type Kelvin int64 // or float64

func CelsiusToKelvin(v float64) Kelvin
func KelvinToCelsius(v Kelvin) float64

func FahrenheitToKelvin(v float64) Kelvin
func KelvinToFahrenheit(v Kelvin) float64
