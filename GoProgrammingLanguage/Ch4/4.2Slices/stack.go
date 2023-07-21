// implementation of stack using slice
package main

import "fmt"

func main() {
	fmt.Println("Stack")
	var employee []string = []string{"Mary", "John", "Paul"}

	// implement stack using slices

	// Add Stack
	fmt.Println(employee)
	employee = addStack(employee, "Kate")

	fmt.Println(employee)

	// Remove Stack
	employee = removeStack(employee)
	fmt.Println(employee)

	// slices is great that you can modify the value
	emptyInitial(employee)
	fmt.Println(employee)

	// assign initial as Bob
	assignInitial(employee)
	fmt.Println(employee)

	// remove index at 1 "john"
	employee = removeIndex(employee, 1)
	fmt.Println(employee)
}

// append by default add the last element
// utilises the built-in append func
func addStack(stack []string, s string) []string {
	var newstack []string
	newstack = append(stack, s)
	return newstack
}

func removeStack(stack []string) []string {
	if len(stack) == 0 {
		return stack
	}
	stack = stack[:len(stack)-1] // len(stack) is not inclusive of the last element. stack[:len(stack)] returns the last element
	return stack
}

func emptyInitial(stack []string) {
	stack[0] = ""
}

func assignInitial(stack []string) {
	stack[0] = "Bob"
}

// remove at specific index
func removeIndex(stack []string, index int) []string {
	// ["a", "b", "c"]
	if index >= len(stack) {
		fmt.Println("Index Out of Bounds")
	}
	copy(stack[index:], stack[index+1:])
	// remove last element :len(stack)-1
	// operation is not inclusive
	// so will decrease the length of the slice by 1
	return stack[:len(stack)-1]

}
