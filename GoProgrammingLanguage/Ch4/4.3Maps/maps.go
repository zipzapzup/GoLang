package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps - HashMaps")

	// two way to define maps
	ages := make(map[string]int) // via make is empty
	ages2 := map[string]int{     //map literal
		"John":  30,
		"Bob":   16,
		"Peter": 18,
	}

	fmt.Println("Map via Make:", ages, "\nMap Literal", ages2)

	ages2["Sarah"] = 20
	ages2["London"] = 30
	delete(ages2, "Bob") // delete remove the key
	delete(ages2, "ZZ")

	fmt.Println("Map via Make:", ages, "\nMap Literal", ages2)

	fmt.Println(ages2["Z"]) // Z does not exist and it will return its zero value 0

	ages2["Bob"]++ // adding it will assign Bob
	fmt.Println("Map via Make:", ages, "\nMap Literal", ages2)

	ages2["Alicia"] = 19

	// Loop through maps

	fmt.Println("\nLoop through maps")
	for name, ages := range ages2 {
		fmt.Printf("Key: %s, Value: %d \n", name, ages)
	}

	sortedKey := make([]string, 0, len(ages2))
	for key := range ages2 {
		sortedKey = append(sortedKey, key)
	}
	sort.Strings(sortedKey)
	fmt.Println(sortedKey, len(ages2))

	// Map Set Operation O(1)
	ages2["Lucky"] = 30

	// Map Access Operation O(1)
	fmt.Println(ages2["Lucky"])

	// Map Check Operation O(1)
	age, ok := ages2["Nonexistant"]
	if !ok {
		fmt.Println(ok) // if not okay, do the following operation
	}
	fmt.Println(age)

	// spin two map literals

	var map1 map[string]int = map[string]int{
		"John": 10,
		"Bob":  12,
		"Pete": 8,
	}

	var map2 map[string]int = map[string]int{
		"John":   10,
		"Bob":    12,
		"Pete":   8,
		"POPEYE": 1,
	}

	fmt.Println("Run Function Equal:", equal(map1, map2))

	ages["ABC"] = 1
	fmt.Println(ages)

}

// run function equal check between two map
// O(N) Time Complexity
// O(1) Space Complexity
func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for item := range a {
		_, ok := b[item]
		if !ok {
			return false // if does not exist, return false
		}
		if a[item] != b[item] {
			return false
		}
	}
	return true

}
