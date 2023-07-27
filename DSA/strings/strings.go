package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Strings")

	a := []string{"ab", "ba"}
	fmt.Println(a)

	joined := strings.Join(a, "")
	fmt.Println(joined)
}
