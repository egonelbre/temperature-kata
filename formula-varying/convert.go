package temp

type Temperature struct {
	Value float64
	Unit  Unit
}

type Unit struct {
	// Kelvin = (Temp - C) / B - A
	// Temp = (Kelvin + A) * B + C
	A, B, C float64
}

func (unit *Unit) Kelvin(v float64) Kelvin { return Kelvin((v-unit.C)/unit.B - unit.A) }
func (unit *Unit) Value(v Kelvin) float64  { return (float64(v)+unit.A)*unit.B + unit.C }

var (
	// TODO: these are relative to celsius
	Celsius    = Unit{0, 1, 0}
	Fahrenheit = Unit{0, 1.8, 32}
	// Kelvin     = Unit{273.15, 1, 0}
	Rankine = Unit{273.15, 1.8, 0}
	Delisle = Unit{-100, -1.5, 0}
	Newton  = Unit{0, 0.33, 0}
	Reaumur = Unit{0, 0.8, 0}
	Romer   = Unit{0, 0.525, 7.5}
)
