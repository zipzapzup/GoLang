package main

import "fmt"

func main() {
	var p = f()
	fmt.Println(p)
	fmt.Println(p)
	// p is the memory address of f() as its returning the &v
	// in the function it is returning of a pointer type

	// now we retrieve the value of the pointer
	// p = 1
	fmt.Println(*p)

	// in here we change the pointer value of p to 20
	// p = 20
	*p = 20
	fmt.Println(*p)

	update(p)
	fmt.Println(*p)

	// in here v is 1
	// to update v you cannot specify just v, but the memory address of v
	v := 1
	update(&v)
	fmt.Println("V is: ", v) // v = 2

	// in here z is still 17 even though the function tries to update the value
	// no assignment is done
	var z = 17
	fmt.Println(addone(z), z)

}

// in this f function below, the f() will return the address of a local variable.
// local variable v will remain in existence after the call has returned.
// multiple call to p above will show the same memory address
func f() *int {
	v := 1
	return &v
}

// since we are passing a pointer
// we can use this function to update the variable indirectly
func update(ptr *int) int {
	*ptr++
	return *ptr
}

// note that in here, when we don't pass the pointer as input
// the value of x will not be updated
// the return value will be x + 1, but value of x is not updated.
func addone(x int) int {
	x = x + 1
	return x
}
