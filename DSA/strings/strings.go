<<<<<<< HEAD
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings")

	// strings.Join([]string, sep) is O(N) ST Operation
	a := []string{"ab", "ba"}
	fmt.Println(a)

	joined := strings.Join(a, "")
	fmt.Println(joined)

	fmt.Println(5381 * 33)

	z := "hello"
	for i, v := range z {
		fmt.Println(i, v, z[i])
		if v == rune(' ') {
			fmt.Println("A space here")
		}
	}

	s1 := "apple pie 你好吧"
	s2 := "i 你好吧"

	fmt.Println(s1, s2)
	fmt.Println(checkStrings(s1, s2))

	stringbyte("Hello")

}

// check strings through loop
func checkStrings(s, sub string) bool {
	// to compare strings in golang
	// you need to convert them to rune
	// r is []rune type
	r := []rune(sub)

	// looping through strings give you runes
	// UTF-8 value
	// O(1) Space Complexity since it accessed the runes value
	for _, v := range s {
		if v == r[0] {
			return true
		}
	}
	return false
}

// check strings through function
func checkStringsv2(s, sub string) bool {
	// Contains will go through each character in strings
	// O(n) Time Complexity
	return strings.Contains(s, sub)
}

func stringLiteral() string {
	return `Mary had a little lambs ?`
}

func stringbyte(s string) {
	// index of string is accessible
	// however it prints the bytes
	fmt.Println("Running stringbyte: hello")
	for index, value := range s {
		fmt.Println(index, value)
	}

	// len(s) is also accessible
	fmt.Println("length of s:", len(s))
}
=======
package main

import "fmt"

func main() {
	fmt.Println("Strings")

	s := "Mary had a little lamb"

	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			fmt.Println(i, "there is a space")
		} else {
			fmt.Println(string(s[i]))
		}
	}
}
>>>>>>> 3a1ba7a (integer)
