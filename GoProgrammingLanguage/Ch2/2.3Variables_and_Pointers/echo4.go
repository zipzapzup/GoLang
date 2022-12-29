package main

import (
	"flag"
	"fmt"
	"strings"
)

// n is a true or false
// flag Bool takes 3 arguments,
// 1 - the name it watch out for in the argument
// 2 - the default
// 3 - the help print message
// flag bool returns Pointer of *Bool

// both n and sep return Pointer type
// n = *Bool
// sep = *String

var n = flag.Bool("n", false, "takes off trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	fmt.Println(" Echo version 4 ")

	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
