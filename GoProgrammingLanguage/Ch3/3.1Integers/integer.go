package main

import "fmt"

func main() {
	var x int = 20
	fmt.Printf("%v", x)

	var y int8 = 127
	fmt.Printf("\n%v\n", y)

	// example of overflow
	// var u uint8 = 255
	// fmt.Println(u, u+1, u*u)
	// use math/big packages

	fmt.Printf("X: %b, %v\n", x, x)
	x = x << 1
	fmt.Printf("Operator: BIT SHIFT to left by 1: X << 1\n")
	fmt.Printf("X: %b, %v\n", x, x)

	// Binary Operator: OR adds it by 1
	fmt.Printf("Operator: BIT OR: X | 1\n")
	x = x | 1
	fmt.Printf("X: %b, %v\n", x, x)

	fmt.Printf("Operator: BIT AND: X & 1\n")
	x = x & 1
	fmt.Printf("X: %b, %v\n", x, x)

	var f float64 = 3.23123123123231
	fmt.Printf("\nFloat 64: %3.4f", f)
	// %.4f means that we are limiting the precision point to 4 decimal places
	// %2.4f means that we are limiting the width to 2 and 4 precision point

}
