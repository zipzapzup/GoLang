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

	a := 126
	fmt.Println(126 % 10)
	a = a / 10
	fmt.Println(a)
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

func splitdigit(num int) int {
	list := []int{}

	if num > 0 {

	}

}
