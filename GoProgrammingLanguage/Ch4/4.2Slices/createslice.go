package main

import "fmt"

func main() {
	fmt.Println("Create Slices")

	var slice1 []int           //initialise slice
	slice1 = make([]int, 4, 8) //assign slice make([]T, len, cap)
	fmt.Println(slice1)
}
