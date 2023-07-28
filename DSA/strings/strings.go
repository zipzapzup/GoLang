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
	}
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
