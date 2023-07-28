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
	for _, i := range z {
		fmt.Println(int(i), i)
	}
}
