package temp

type Celsius float64
type Kelvin float64
type Fahrenheit float64

func (v Celsius) Fahrenheit() Fahrenheit
func (v Celsius) Kelvin() Kelvin

func (v Fahrenheit) Celsius() Celsius
func (v Fahrenheit) Kelvin() Kelvin

func (v Kelvin) Celsius() Celsius
func (v Kelvin) Fahrenheit() Fahrenheit
