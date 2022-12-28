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
