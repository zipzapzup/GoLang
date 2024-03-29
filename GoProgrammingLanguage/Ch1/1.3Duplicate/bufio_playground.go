// go run bufio_playground < newline.txt
// bufio playground reads standard input
// and return a count on the number of words
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	// establish a new map
	count := make(map[string]int)

	// input source is os.stdin
	// NewScanner establishes a new scan object
	input := bufio.NewScanner(os.Stdin)
	fmt.Println(input)

	// Scanner interface has a scan() function
	// that allows it to loop return true until EOF
	// looping through the scan
	// allow you to convert input via text or bytes

	fmt.Println("Looping through Scanner")
	for input.Scan() {
		line := input.Text()
		// BigO Time: O(n)
		// BigO Space: O(n)
		// where n refers to the number of words
		// 1 * O(n) operations for going to the line
		// 2 * O(n) operations for converting to bytes plus string

		fmt.Println(line)
		words := strings.Split(line, " ")
		// 1 * O(n) operation on each line converting to word

		for _, i := range words {
			if i == "" {
				continue
			}
			count[i] = count[i] + 1
			// 1 * O(n) operation for counting
		}
	}

	fmt.Println("001. Count How Many Words in the Input")
	for key, value := range count {
		fmt.Printf("%d\t %s\n", value, key)

	}

	// Sort the keys and use the previous maps
	fmt.Println("002. Printing Sorted")
	sortedKeys := []string{}
	sortedKeys = sortMap(count)
	for _, value := range sortedKeys {
		fmt.Printf("%d \t %s\n", count[value], value)
	}

}

func sortMap(mapInput map[string]int) []string {
	stringList := []string{}
	for key, _ := range mapInput {
		stringList = append(stringList, key)
	}

	// sort Strings
	sort.Strings(stringList)

	return stringList
}
