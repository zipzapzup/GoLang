package main

import (
	"fmt"
	"os"
)

func main() {

	var s string
	for _, arg := range os.Args[1:] {
		s += " " + arg
	}
	fmt.Println(s)
	// Improved version of loop
	// for loop with range omits two values, index and value.
	// since value is not used, we use _

}

// BigO O(n) Time Complexity
// BigO O(n) Space Complexity
