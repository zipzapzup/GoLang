package main

import "fmt"

func main() {
	fmt.Println("Hello GoLang")

	var x int = 10
	var y float64 = 2.0
	var z float32 = 3.2
	fmt.Println(x, y, z)

	var s string = "A brand new World"
	fmt.Println(s)

	fmt.Println("Basic Primitive Types in Golang:")

	var a bool = true
	fmt.Println("1. Boolean: ", a)
	var b string = "A string"
	fmt.Println("2. String: ", b)
	var c int = 50
	fmt.Println("3. Int: ", c)
	var d uint16 = 5 // 0 to 65535
	fmt.Println("3. Uint (Only Positives): ", d)
	var e byte = 20
	fmt.Println("4. Byte (Uint8): ", e) // 0 to 255 only
	var f float32 = 3.5123
	fmt.Println("5. Float32: ", f)
	var g complex64 = 12 + 23i
	fmt.Println("6. Complex64 or Imaginary Number: ", g)
	var h rune = 'a'
	fmt.Println("7. Rune: ", h) // Rune Prints the UNICODE Character of the character

}
