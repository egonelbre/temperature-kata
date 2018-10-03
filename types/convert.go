package temp

type Celsius float64
type Kelvin float64
type Fahrenheit float64

func CelsiusToFahrenheit(v Celsius) Fahrenheit
func FahrenheitToCelsius(v Fahrenheit) Celsius

func CelsiusToKelvin(v Celsius) Kelvin
func KelvinToCelsius(v Kelvin) Celsius

func FahrenheitToKelvin(v Fahrenheit) Kelvin
func KelvinToFahrenheit(v Kelvin) Fahrenheit
