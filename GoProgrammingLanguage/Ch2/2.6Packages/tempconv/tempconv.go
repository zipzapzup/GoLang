package tempconv

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZero Celcius = -273.13
	FreezingC    Celcius = 0
	BoilingC     Celcius = 100
)

// Celcius types method
func (c Celcius) String() string {
	return fmt.Sprintf("%g °C", c)

}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}
