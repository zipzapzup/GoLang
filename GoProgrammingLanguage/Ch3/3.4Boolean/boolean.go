package main

import "fmt"

func main() {
	fmt.Println("Boolean")
	var a bool = true
	fmt.Printf("booltoint conversion: %d\n", booltoint(a))
	var num int = 0
	var num1 int = 1
	fmt.Printf("inttobool conversion: %t\n", inttobool(num))
	fmt.Printf("inttobool conversion: %t\n", inttobool(num1))

}

func booltoint(b bool) int {

	if b {
		return 1
	}
	return 0
}

func inttobool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}
