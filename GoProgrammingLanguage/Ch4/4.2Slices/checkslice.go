package main

import (
	"fmt"
	"sort"
)

func main() {
	var s1 []string
	s1 = []string{"John", "BOB", "Alex", "Alexandria", "Lucy"}

	fmt.Println(s1)

	s1 = append(s1, "Peter")
	fmt.Println(s1)

	sort.Strings(s1)
	fmt.Println(s1)
	s1 = append(s1, "Peter")

	fmt.Println(s1)

}
