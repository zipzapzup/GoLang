package main

import (
	"example/user/hello/morestrings"
	"example/user/hello/morestrings2"
	"fmt"
)

func main() {
	fmt.Println("hello, world.")

	fmt.Println(morestrings.ReverseRunes("!oG, olleh"))

	fmt.Println(morestrings.Adding(5, 6))
	fmt.Println(morestrings2.More())

	fmt.Println(three)

}
