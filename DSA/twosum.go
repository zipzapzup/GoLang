package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	input := []int{2, 7, 11, 8, 15}

	fmt.Println(twoSum(input, 9))
	// for i := 0; i < len(input); i++ {
	// 	for j := 0; j < len(input); j++ {
	// 		if i == j {
	// 			continue
	// 		}
	// 		fmt.Println("i var", i)
	// 		fmt.Println("j var", j)
	// 	}
	// }

	fmt.Println(twoSumFast(input, 9))

}

// O(N^2) Time Complexity
// O(1) Space Complexity
func twoSum(num []int, target int) []int {
	result := make([]int, 0, 2)
	sum := 0
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if i == j {
				continue
			}
			sum = num[i] + num[j]
			if sum == target {
				result = append(result, i, j)
				return result
			}
		}
	}

	return result
}

// Goal is faster
// O(N) Space Complexity
// O(N) Time Complxity
// Algorithmn uses a hashmaps
func twoSumFast(nums []int, target int) []int {
	hashmap := make(map[int]int) // O(N) Space Complexity for hashmap
	result := make([]int, 0, 2)  // O(1) Space Complexity

	// create hashmap of index:value
	for i := range nums {
		hashmap[nums[i]] = i
		// O(N) Operation for Time Complexity
	}

	// loop through nums
	// check if the complement exist
	for i := range nums {
		// O(N) Operation for Time Complexity
		complement := target - nums[i]
		value, ok := hashmap[complement]
		if ok {
			result = append(result, i, value)
			return result
		}
	}
	return result

}
