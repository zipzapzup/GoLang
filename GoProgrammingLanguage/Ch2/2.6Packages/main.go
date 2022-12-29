package main

import (
	"fmt"

	"2.6Packages/tempconv"
)

func main() {
	fmt.Printf("Cold! %v\n", tempconv.AbsoluteZero)
	fmt.Printf("Convertion to Fahrenheit: %v\n", tempconv.CToF(tempconv.BoilingC))
}
