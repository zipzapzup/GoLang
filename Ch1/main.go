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

	// base 16 - hexadecimal
	// 14 hexa is 20
	var e int = 0x14
	fmt.Println("Base 10:", a+b, "Binary: ", c, "Octal: ", d, "Hexa:", e)

	// float - decimal values
	// Has 2 values: float32 float64
	var f float32 = -3.2342323
	fmt.Println(f)

	// rune - unicode value (character set mapping)
	// represented via single quotes ''
	var g rune = '1'
	var h rune = 'h'
	fmt.Println(g, h)

	// strings
	// use backquote to escape the strings - \n is newline \" will show quotes
	var i string = "Greetings and \n\"Salutations\""
	fmt.Println(i)

	// strings - raw strings literal
	// use backquotes to represent the raw strings ``
	var j string = `" And in Here you can get \n all Raw\n Strings`
	fmt.Println(j)
}
