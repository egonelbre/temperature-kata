package temp

type Temperature float64 // or fixnum

func (temp *Temperature) SetFahrenheit(v float64)
func (temp Temperature) Fahrenheit() float64

func (temp *Temperature) SetCelsius(v float64)
func (temp Temperature) Celsius() float64

func (temp *Temperature) SetKelvin(v float64)
func (temp Temperature) Kelvin() float64
