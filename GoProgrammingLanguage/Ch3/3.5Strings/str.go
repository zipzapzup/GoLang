package main

import "fmt"

func main() {
	fmt.Println("Strings")

	var s string = "A sentence."
	fmt.Printf(s)

	fmt.Println("\nlength of s string:", len(s))
	fmt.Println(s[0])
	// note this print the bytes

	// loop through the string and print
	fmt.Println("Loop through print it")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println(s[2:10])
	// substring create new string

	var z string = "AB"
	var z1 string = "世界"
	fmt.Println(z, z1)
	fmt.Println("Comparison")
	fmt.Printf("ASCII char length: %d, Non ASCII char length: %d", len(z), len(z1))

	fmt.Println("\n", z1[1])

	fmt.Println("Printing Unicode Characters")
	var s2 string = "Hello, 世界"
	for i, r := range s2 {
		fmt.Printf("Loop: %d\t %q\n", i, r)
	}

	var count int = 0
	for char := range s2 {
		fmt.Println(char)
		count++
	}
	fmt.Println("Characters are: ", count)

	// You can convert a string to a rune.

	var s3 string = "Hello, こんにちは"
	r1 := []rune(s3)

	// when printing the value, it prints the unicode value of it
	// encapsulate it in string() to convert it to string
	//
	fmt.Println("RUNE Operations int32")
	for i, value := range r1 {
		fmt.Printf("%d. %v and %q\n", i, value, value)
	}
}
