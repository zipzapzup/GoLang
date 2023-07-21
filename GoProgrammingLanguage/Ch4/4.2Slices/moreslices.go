package main

import "fmt"

func main() {
	fmt.Println("More Slices")
	var s []int = []int{1, 4, 12, 23, 4, 5, 5, 99}
	// len 8
	// cap 8
	// slices string
	fmt.Println(len(s))

	// print initial element
	fmt.Println(s[0])

	// print 2nd element
	fmt.Println(s[1])

	// print all
	fmt.Println(s[0:])

	fmt.Println(s[1:])

	// len is 8 and s[7] prints last
	fmt.Println(s[7])

	// len(s)-1 is the last element
	fmt.Println(s[len(s)-1])

	// s[len(s):] prints all elements since len(s) is not inclusive
	// to print excluding the last element do len(s)-1
	fmt.Println("s[:8]", s[:])
	fmt.Println(s[:6])
}
