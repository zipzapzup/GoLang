package main

import (
	"fmt"
	"strconv"
)

type House struct {
	name   string            // O(1)
	people map[string]person // O(N)
	count  int               //O(1)
}

type person struct {
	name string //O(1)
	age  int    //O(1)
} // Space O(2+ N(2)) = O(2 + 2N)  // O(N) Space Complexity

func main() {

	var p1, p2, p3 person
	p1 = person{name: "John", age: 10}

	p2 = person{name: "Sarah", age: 20}
	p3 = person{name: "Doe", age: 100}

	bigHouse := new(House) // initialise
	fmt.Println(bigHouse)

	bigHouse.storeHouse("House of LV", p1)
	fmt.Println(bigHouse)
	bigHouse.storeHouse("House of LV", p2)

	bigHouse.storeHouse("House of LV", p3)

	fmt.Println(bigHouse)

	fmt.Println("Running Second Operation")
	fmt.Println(bigHouse.listEveryone())

	fmt.Println("How many people in the house?")
	fmt.Println(bigHouse.countPeople())

	fmt.Println("Checking People")
	value, err := bigHouse.checkPeople("Sarah", 20)
	fmt.Println(value, err)

	fmt.Println("Checking 2nd People")
	value, err = bigHouse.checkPeople("John", 11)
	fmt.Println("Printing .. ", value, err)

	fmt.Println(addNumber(2, 2))

}

func (h *House) checkPeople(name string, age int) (bool, error) {
	var error error
	var result bool
	var checkPerson person = person{name: name, age: age}

	for _, value := range h.people {
		if value == checkPerson {
			fmt.Println("Found! Person is in the house", value)
			result = true
		}
	}
	if !result {
		error = fmt.Errorf("person not found")
	}
	return result, error

}

func (h *House) countPeople() int {
	return len(h.people)
}

func (h *House) listEveryone() []person {
	var persons []person
	persons = make([]person, 0, len(h.people))

	fmt.Println("Listing Everyone who is in the house")
	for key, value := range h.people {
		fmt.Println(key, value.name, value.age)
		persons = append(persons, value)
	}

	return persons

}

func (h *House) storeHouse(name string, p person) string {
	var parameterId string
	h.name = name
	// allocations step
	// allocate the person struct
	// fmt.Println("Check Before Allocation", h.people == nil)
	if h.people == nil {
		h.people = make(map[string]person)
	}
	// fmt.Println("Check After Allocation", h.people == nil)

	// generate parameterId as a key
	parameterId = strconv.Itoa(h.count)
	h.count++

	// alocate the person to the map
	// parameterId is a reference to the person
	h.people[parameterId] = p

	return parameterId
}

func addNumber(x, y int) int {
	return x + y
}
