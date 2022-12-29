package main

import "fmt"

func main() {
	fmt.Println("Lessons About Pointers")

	// Declare x of an integer type with 0 value
	var x int = 0
	// ptr here refers to the memory address of x using &x
	ptr := &x

	// in here we print the "ptr", the pointer that holds the memory address
	// yields 0x140000aa008
	fmt.Println("Pointer: ", ptr, "AND Value: ", *ptr)

	// now if we want to change the value
	*ptr = 10
	fmt.Println("Pointer: ", ptr, "AND Value: ", *ptr)

	// Lessons About Pointers
	// Pointer:  0x14000122008 AND Value:  0
	// Pointer:  0x14000122008 AND Value:  10
	// Note in the above expressions, we changed the value of the pointer *ptr to 10

	// Any variables is an addressable values. This means that we can use & memory address operations to any Var

	// Pointers are also comparable

	var z, y int
	fmt.Println("Comparing Memory Address of Z and Y, '&z == &y': ", &z == &y)
	fmt.Println("Comparing Memory Address of Z and Z, '&z == &z': ", &z == &z)
	fmt.Println("Comparing Memory Address of Z and Nil, '&z == nil' it is not nil since Z has a memory address: ", &z == nil)

	var x1 *int
	var x2 int
	fmt.Println(x1, x2)
}
