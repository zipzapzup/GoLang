package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	// fixed length of element sequence of the same type

	// array of 3 strings
	var arr [3]string

	// array2 of 3 ints
	var arr2 [3]int

	// printing it generates the zero value
	fmt.Println(arr[0], arr2[0])

	// show length
	fmt.Println(len(arr), len(arr2))

	// last element
	fmt.Println(arr[len(arr)-1])

	// loop via println

	fmt.Println("Looping through Array2 INT")
	for i, v := range arr2 {
		fmt.Println("Index: ", i, "Value: ", v)
	}

	// loop via prinf
	fmt.Printf("Looping through Array1 STRING via printf")
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %v\n", i, v)
	}

	// array literal allow you to initialise the value along the size
	var arr11 [3]int = [3]int{9, 7, 5}
	fmt.Println("Arr11 value: ", arr11[0])

	arr12 := [...]int{1, 2, 3, 4, 5, 5}
	fmt.Println("Arr12 value using array literals and shortform: ", arr12[2])

	// Array of Key value Pair

	// type Weekdays string

	// const (
	// 	Mon Weekdays = iota
	// 	Tue
	// 	Wed
	// 	Thu
	// 	Fri
	// )

	workday := [...]string{1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}

	fmt.Println(workday, workday[5])

	// Short Hand Array declaration
	r := [...]int{99: -19}
	fmt.Println(r)

	// Arrays can be compared if the type are comparable and of same length
	arrCompare1 := [3]int{1, 2, 3}
	arrCompare2 := [3]int{1, 2, 3}
	arrCompare3 := [3]int{1, 2, 5}

	fmt.Println(arrCompare1 == arrCompare2) // true since the same
	fmt.Println(arrCompare1 != arrCompare3) //  true since not the same

	arrCompare1[2] = 10
	fmt.Println("arrCompare1", arrCompare1) // arrays are modifieable

	// however arrays size cannot be changed
	// arrays does not have append or remove function to remove or add extra elements

}

// function to empty an array to 0
// use a pointer so the array elements is visible to the caller
// more efficient than creating multiple arrays
// as go will develop a copy of the array
func zeroArray(ptr *[32]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}

// improved version
// set it to empty array immediately
// pointer efficiently mutate caller's variable
func zeroArray2(ptr *[32]int) {
	*ptr = [32]int{}
}
