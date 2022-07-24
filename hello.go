package main

import "fmt"

func main() {

	// integer literals

	// base 10
	var a int = 10
	var b int = 20

	// base 2 - binary
	var c int = 0b01010

	// base 8 - octal
	// 024 octal is 20
	var d int = 024

	fmt.Println(a+b, "Binary: ", c, "Octal: ", d)

	fmt.Println("Hello World")

}
