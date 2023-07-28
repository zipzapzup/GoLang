package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func main() {
	fmt.Println("Strings")

	// strings.Join([]string, sep) is O(N) ST Operation
	a := []string{"ab", "ba"}
	fmt.Println(a)

	joined := strings.Join(a, "")
	fmt.Println(joined)

	fmt.Println(5381 * 33)

	z := "hello in a new world"
	r1 := []rune(" ")
	for i, v := range z {
		fmt.Println(i, v, z[i])
		if v == r1[0] {
			fmt.Println("A space here")
		}
	}

	s1 := "apple pie 你好吧"
	s2 := "i 你好吧"

	fmt.Println(s1, s2)
	fmt.Println(checkStrings(s1, s2))

	stringbyte("Hello")

	c1 := 'c'
	fmt.Println(reflect.TypeOf(c1))

	fmt.Println("compareStrings :")
	fmt.Println(compareStrings("ABC", "ABC"))

	fmt.Println(s1 == s2)

	checkStringsv2("asd", "d")

	stringLiteral()

	countString("helloo", "o")

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[:0], arr[:1], arr[2:3])

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

func TestCheckStringsAnotherVersion(t *testing.T) {
	var x = []struct {
		input1, input2 string
		result         bool
	}{
		{
			input1: "hello there world",
			input2: "o",
			result: true,
		},
		{
			input1: "arrivederci",
			input2: "i",
			result: true,
		},
	}

	for _, test := range x {
		if checkStrings(test.input1, test.input2) != test.result {
			t.Error(`checkStrings( ,) !=,v`, test.input1, test.input2, test.result)
		}
	}

}

// check strings through function
func checkStringsv2(s, sub string) bool {
	// Contains will go through each character in strings
	// O(n) Time Complexity
	return strings.Contains(s, sub)
}

func compareStrings(a, b string) bool {
	return a == b
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

func countString(s, c string) int {
	return strings.Count(s, c)
}
