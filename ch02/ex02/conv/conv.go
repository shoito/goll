package conv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Feet) String() string  { return fmt.Sprintf("%gFt", f) }
func (m Meter) String() string { return fmt.Sprintf("%gM", m) }

func (p Pound) String() string    { return fmt.Sprintf("%gLb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }

func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.4536) }

func KToP(k Kilogram) Pound { return Pound(k / 0.4536) }
