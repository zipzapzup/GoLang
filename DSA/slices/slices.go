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

	arr3 = insertv2(arr3, 1, 100)
	fmt.Println(arr3)

	arr5 := []int{}

	fmt.Println(&arr5)

	arr5 = append(arr5, 1)
	// memory address different
	fmt.Println(&arr5[0])

	// memory address different suggests O(Log(N) Operation for Memory?)
	arr5 = append(arr5, 1)
	fmt.Println(&arr5[0])

}

// insert func to insert a slice at specific index
//
//	1, -10, 7, 10000, 8
//
// O(N) Time Complexity
// O(N) Space Complexity
func insert(arr []int, index, value int) []int {
	length := len(arr)

	// operation to extend the array O(n)
	arr = append(arr[:index+1], arr[index:]...)
	arr[index] = value

	return arr[:length]
}

// O(N) Time Complexity
// O(1) Space Complexity
func insertv2(arr []int, index, value int) []int {
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = value
	return arr
}
