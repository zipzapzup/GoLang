package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("leetcode style")

	fmt.Println(lengthOfLastWord("Mary had a little lamb "))
	fmt.Println(fizzBuzz2(15))
}

// leetcode 58
func lengthOfLastWord(s string) int {
	var n int = len(s)
	var count int = 0

	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				break
			}
			continue
		} else {
			count++
		}
	}
	return count
}

// leetcode 412
// BigO: T:O(N^2) ,S:O(N)
func fizzBuzz(n int) []string {
	var answer = []string{}
	// go1.22: for index := range n
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if i%3 == 0 {
			answer = append(answer, "Fizz")
		} else if i%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(i))
		}
	}
	return answer
}

// leetcode 412 # iteration 2
// BigO(ST) T:O(N) ,S:O(N)
func fizzBuzz2(n int) []string {
	var answer = make([]string, n)
	// create slice of string of n length and cap.

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			answer[i-1] = "Fizz"
		} else if i%5 == 0 {
			answer[i-1] = "Buzz"
		} else {
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}
