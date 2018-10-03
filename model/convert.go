package temp

type Celsius float64
type Kelvin float64
type Fahrenheit float64

func (v Celsius) Kelvin() Kelvin
func (v Fahrenheit) Kelvin() Kelvin
func (v Kelvin) Kelvin() Kelvin

type Temperature interface {
	Kelvin() Kelvin
}

type Model interface {
	Convert(t Temperature) Temperature
}

var (
	CelsiusModel    = ModelFunc(celsiusModel)
	FahrenheitModel = ModelFunc(fahrenheitModel)
	KelvinModel     = ModelFunc(kelvinModel)
)

type ModelFunc func(Temperature) Temperature

func (m ModelFunc) Convert(t Temperature) Temperature { return m(t) }

func celsiusModel(t Temperature) Temperature {
	kelvin := t.Kelvin()
	return Celsius(kelvin)
}

func fahrenheitModel(t Temperature) Temperature {
	kelvin := t.Kelvin()
	return Fahrenheit(kelvin)
}

func kelvinModel(t Temperature) Temperature {
	kelvin := t.Kelvin()
	return Kelvin(kelvin)
}
