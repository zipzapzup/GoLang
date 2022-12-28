package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting Program")

	for index, value := range os.Args {
		fmt.Println(index, value)
	}

	fmt.Println("Second Loop Iteration")

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])

	}
}
