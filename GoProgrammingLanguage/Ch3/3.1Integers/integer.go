package main

import "fmt"

func main() {
	var x int = 20
	fmt.Printf("%v", x)

	var y int8 = 130
	fmt.Printf("\n%v\n", y)

	// example of overflow
	// var u uint8 = 255
	// fmt.Println(u, u+1, u*u)
	// use math/big packages
}
