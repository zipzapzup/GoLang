package main

import "fmt"

// package declaration
const boilingPoint = 212.0
const Boiling = 212.0

func main() {
	// function scope declaration within main function
	var f = boilingPoint
	var c = (f - 32) * 5 / 9

	// details of fmt.Printf see: https://pkg.go.dev/fmt#example-package-Formats
	// %v = default format value
	// %t = boolean true or false
	// %b = base 2
	// %d = base 10
	// %.2f = float points with 2 precision on default width
	// %9.2f = width 9 with 2 precision point.

	fmt.Printf("boiling point = %v, %v \n", f, c)

	any()

}

func any() {
	fmt.Printf("Execution from any function: %v", boilingPoint)
}
