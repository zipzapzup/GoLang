package main

import "fmt"

func main() {

	// Declares that x is a variable of int type
	var x int = 20
	fmt.Println(x, &x)

	var y int = 5

	fmt.Println("value of x and memory address:", x, &x)
	fmt.Println("value of y and mmeory address:", y, &y)
	// Delcares ptr is a variable of pointer int,
	// which points to the memory address of x
	var ptr *int = &x
	*ptr = 40
	fmt.Println(x, &x)
	fmt.Println("Pointer Value and Address:", ptr, &ptr)

	ptr = &y
	fmt.Println("Pointer Value and Address:", ptr, &ptr)

	var ptr2 *int = &y
	fmt.Println("Pointer 2: ", ptr2)

}
