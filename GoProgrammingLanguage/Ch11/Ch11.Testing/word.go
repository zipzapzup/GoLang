package word

// func main() {
// 	fmt.Println("Testing")

// 	fmt.Println(addNumber(5, 5))

// 	// palindrome
// 	fmt.Println("Palindrome")
// 	fmt.Println(isPalindrome("abggba"))
// 	fmt.Println(isPalindrome("こんんこ"))

// 	fmt.Println(isPalindromv2("こんんこ"))

// }

func addNumber(x, y int) int {
	return x + y
}

// only can process ASCII code
func isPalindrome(s string) bool {
	// unicode
	// loop through the index
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// improved palindrome which can process all unicode and characters
// Time BigO 2N + 1 or  BigO(N)
// Space Complexity O(1) since we are not creating additional strings. Operations is in Unicode / Runes / int32
func isPalindromev2(s string) bool {

	// fmt.Println("Start Palindrom v2")

	// convert to slice O(n)
	stringSlice := []int32(s)
	length := len(stringSlice) //O(1)

	for index := range stringSlice { // O(n)
		if stringSlice[index] != stringSlice[length-1-index] {
			return false
		}

	}
	return true
}
