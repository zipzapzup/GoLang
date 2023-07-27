package main

import "fmt"

func main() {
	fmt.Println("Slices")

	// create slice of 0 len and 5 cap
	arr1 := make([]int, 0, 5)
	fmt.Println(arr1)

	// slices of 5 len and 5 cap
	// difference is the zero value
	arr2 := make([]int, 5)
	fmt.Println(arr2)

	// slices
	arr3 := []int{}
	fmt.Println(arr3)

	// all slices have been initialised
	fmt.Println(arr1 == nil, arr2 == nil, arr3 == nil)

	// arr3 = []
	// append is O(n) operation at worst when cap has exceeded because copy
	arr3 = append(arr3, 1, 10, 7, 10000, 8)
	fmt.Println(arr3)

	fmt.Println(arr3[:])
	fmt.Println(arr3[:1])
	fmt.Println(arr3[1:])

	arr3 = insert(arr3, 1, -2)
	fmt.Println(arr3)

}

// insert func to insert a slice at specific index
//
//	1, -10, 7, 10000, 8
func insert(arr []int, index, value int) []int {
	length := len(arr)

	// operation to extend the array O(n)
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value

	return arr[:length]
}
