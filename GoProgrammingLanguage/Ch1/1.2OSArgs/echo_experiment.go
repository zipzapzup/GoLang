package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	timer1 := time.Now()
	fmt.Println("Normal echo")
	echo_normal()
	fmt.Println(time.Since(timer1))

	timer2 := time.Now()
	fmt.Println("Strings echo")
	echo_string()
	fmt.Println(time.Since(timer2))

}

// BigO Time: O(N)
// BigO Space: O(N^2)

func echo_normal() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		sep = " "
		s += os.Args[i] + sep
	}
	fmt.Println(s)
}

// BigO Time: O(N)
// BigO Space: O(N)
func echo_string() {
	s := strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}

// Normal echo
// a b c d e f g h i j k l asdasd sadasdad asdasdasd asdasdsa asdasdas okay longer string how about it
// 22.917µs (ms)
// Strings echo
// a b c d e f g h i j k l asdasd sadasdad asdasdasd asdasdsa asdasdas okay longer string how about it
// 1.958µs (ms)
