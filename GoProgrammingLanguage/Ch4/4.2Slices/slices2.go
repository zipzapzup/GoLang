// slices 2
// playing with slices internals
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Slices 2")

	var s []string = []string{"e", "a", "r", "t", "h"}

	s2 := s[2:]
	fmt.Println(s)
	fmt.Println(s2)

	s2[1] = "l"
	s2[2] = "s"
	// modifying s2 modifies the original slices
	// s = earls
	// s2 creates a new slice that points to s
	// modifying the original value will modify the slice
	//
	// make([]int, len, cap)
	// a short way to create slice
	fmt.Println(s)

	s3 := make([]int, 5, 10)
	fmt.Println(s3)

	start := time.Now()
	s4 := s
	fmt.Println(s4)
	elapsed := time.Since(start)
	fmt.Printf("Assign slice: %s \n", elapsed)

	start = time.Now()
	s5 := s[:]
	fmt.Println(s5)
	elapsed = time.Since(start)
	fmt.Printf("Assign slice [:] %s \n", elapsed)

	// both above show same execution O(N) time

}
