package main

import "fmt"

func main() {
	fmt.Println("Ex: 4.3")

	var arr1 [5]int = [5]int{18, 2, 17, 20, 5}

	fmt.Println(arr1)

	reverseArray(&arr1)
	fmt.Println(arr1)

	zeroArray(&arr1)
	fmt.Println(arr1)

	fmt.Println("Ex: 4.4")
	var slice1 []int = []int{1, 2, 3, 4, 5}

	fmt.Println(slice1)
	fmt.Println("Reverse Slice")
	reverseSlice(slice1)
	fmt.Println(slice1)

	fmt.Println("Zero Slice")
	zeroSlice(slice1)
	fmt.Println(slice1)

	// Array you can assign the index
	// Note it starts from zero
	var element [5]string = [5]string{0: "One", 1: "Two", 2: "Three", 3: "Four", 4: "Five"}
	for i, item := range element {
		fmt.Println(i, item)
	}

	slice1 >= slice1
}

func reverseArray(arr *[5]int) {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func zeroArray(arr *[5]int) {
	for i, _ := range arr {
		arr[i] = 0
	}
}

// rotate slice by n elements
// {1,2,3,4,5}
// rotate by 2 mean
// {3,4,5,1,2}

func zeroSlice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
