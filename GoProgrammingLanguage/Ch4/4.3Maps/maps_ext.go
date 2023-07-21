package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps Still")

	// Try to create a map where the keys are ordered slices

	map1 := make(map[string]int)
	order := make([]string, 0)

	fmt.Println(map1, order)

	// order add the string
	order = add(order, map1, "Bob", 10)
	fmt.Println(map1, order)

	order = add(order, map1, "John", 10)
	order = add(order, map1, "Bob", 20)
	fmt.Println(map1, order)

	// add more user in different format and order
	order = add(order, map1, "Alicia", 20)
	order = add(order, map1, "zaydon", 20)
	order = add(order, map1, "monument", 20)
	order = add(order, map1, "peter", 20)
	order = add(order, map1, "Ally", 20)

	// loop through it
	for _, value := range order {
		fmt.Printf("%s\n", value)
	}

	// initialise map
	var maps2 map[string]int
	fmt.Println(len(maps2))

	if maps2 == nil {
		fmt.Println("Maps2 is Nil")
	}

	// to use maps, assign / allocate
	maps2 = make(map[string]int)
	maps2["BOB"] = 1

	maps2 = map[string]int{
		"Peter": 2,
		"ASD":   2,
	}

	fmt.Println(maps2)
}

func add(s []string, m map[string]int, element string, value int) []string {

	_, ok := m[element]
	if !ok {
		// if key does not exist add the element
		fmt.Println("Operation: Element does not exist, adding them")
		s = append(s, element)
		m[element] = value
	} else {
		// otherwise update the element to the value
		fmt.Println("Operation: Element exist, updating value")
		m[element] = value
	}

	sort.Strings(s)
	return s
}
