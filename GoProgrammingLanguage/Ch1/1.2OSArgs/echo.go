package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i := 1; i < len(os.Args); i++ {
		s = s + " " + os.Args[i]
	}
	fmt.Println(s)
}

// BigO O(n) Space - String Concatenation
// BigO O(n) Time
//
// Important point, the string "s"
// s = s + " " + os.Args[i]
// is assigned a new value
//
// Loop 1 = " args1" |  10 bytes
// Loop 2 = " args1 args2" | 24 bytes
// Loop 3 = " args1 args2 args3" | 32 bytes
//
// Space Complexity: 10 bytes on 1st loop
// Space Complexity: 34 bytes on the 2nd loop
// Space Complexity: 76 bytes on the 3rd loop
//
//
// string will occupy a quadratic space in proportion to the number of args
// s will be garbage collected in due time
