// 合ってるか不安なので後で見直し
// with Kelvin scale
// 0K = -273.15°C
// diff of 1K = diff of 1°C
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100

	ZeroK Kelvin = -273.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c - ZeroK) }
func KToC(k Kelvin) Celsius { return Celsius(k + ZeroK) }

func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k) }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%gK", k) }
