// slices package
// display slices
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Slices ")

	var slice1 []string = []string{"apple", "orange", "cherry", "melon", "grape", "dragonfruit", "banana", "watermelon", "durian", "blueberry"}
	fmt.Println("Original Slices:")
	fmt.Println(slice1)

	// reverse 1 function
	fmt.Println(reverse(slice1))

	// reverse 2 function
	reverse2(slice1[:]) //
	fmt.Println(slice1)

	// interesting observation that the func reverse is faster. Simplicity always better.
	// important part about slices. Passing a slices to a function permits them to modify the elements

}

// reverse slices 1
// one idea is to create new slice
// and append in reverse order
//
// BigO T: O(n) - as the number of element grows, operations grow linearly
// BigO S: O(n) - as the number of elemnt grows, space takes 2(n)
// Total BigO Time: 1N
func reverse(s []string) []string {
	start := time.Now()
	var reversedSlice []string = []string{}

	for i := len(s) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, s[i]) // BigO: O(n)
	}
	elapsed := time.Since(start)
	fmt.Printf("\nFUNC: Reverse Took: %s microseconds\n", elapsed)
	return reversedSlice
}

// reverse slices 2
// another idea is switch the order
// using 2 pointer
// BigO Time: O(n) - execution would take 1/2 N
// BigO Space: O(1)
// Total BigO Time: 4N
// for reverse2 since []string contains a pointer
// modifying the function directly will cause it to be modified
// no need to return another one
func reverse2(s []string) {
	start := time.Now()
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { // BigO Time: 2(N)
		s[i], s[j] = s[j], s[i] // BigO Time: 2(N)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nFUNC: Reverse2 took: %s microseconds\n", elapsed)
}

// reverse3
// try to improve it more
func reverse3(s []string) {

}
