package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World, Functions")

	fmt.Println(hypo(3.0, 4.0))

	fmt.Println(variadicFunc(3, 3, 3, 3, 3))

	deferFunc()

	fmt.Println(fibRecursive(5))

	fmt.Println(inifinity(0))
}

func hypo(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func variadicFunc(x ...int) int {
	sum := 0
	for _, i := range x {
		sum = sum + i
	}
	return sum

}

func deferFunc() {
	defer printLines()

	fmt.Println("Functions Walkthrough")
}

func printLines() {
	fmt.Println("Hello - from deferred")
}

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

func inifinity(n int) int {
	// recover()
	// must be used within defer
	// pass the recover to a variable
	// variable can be used to check if its nil and print the panic value or to perform last resort

	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("Recovered n = 1", a)
		}
	}()

	if n == 0 {
		panic("asd")
		// panic prints a stack error based on the issue it has observed
		// stack trace is available
	}
	return 5 / n
}
