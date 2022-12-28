// Dup prints the text of each line that appears more than one
// in the standard input, preceded by its count
// echo "ASD\n World\n World\n World" | ./dup1

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	count := make(map[string]int)
	// We initiatlise an empty dictionary (map) with the key/value pair of string/int

	input := bufio.NewScanner(os.Stdin)
	// os.Stdin refers to the OS input
	// bufio NewScanner returns the new scanner object

	for input.Scan() {
		count[input.Text()]++

	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
	fmt.Println("Complete")

}
