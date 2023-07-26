package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(absoluteValue(30))
	fmt.Println("OR")
	fmt.Println(math.Abs(float64(30)))

}

// absoluteValue function helps with int
// no built-in func to make absolute value of an int
// math.Abs(float64) exist for float64
func absoluteValue(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
