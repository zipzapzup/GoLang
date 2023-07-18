package main

import "fmt"

func main() {

	var c complex64 = 64 + 2i
	fmt.Printf("%T, %[1]v\n", c)

	fmt.Printf("\nreal number: %v", real(c))
	fmt.Printf("\nimaginary number: %v", imag(c))
}
