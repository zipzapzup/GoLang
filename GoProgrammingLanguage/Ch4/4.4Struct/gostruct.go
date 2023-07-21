package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	Age           int
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	fmt.Println("Struct")

	var Bob Employee // assign bob to an employee
	Bob.ID = 1
	Bob.Age = 10
	Bob.Name = "Bob"
	Bob.Address = "Test Street"
	Bob.Salary = 1

	fmt.Println(Bob)

	Bob.ID = 10 // Modify Bob
	fmt.Println(Bob)

	// since bob is a variable, we can access its variable
	var ptr = &Bob
	*&ptr.Address = "New Address Street"

	fmt.Println(Bob)

	var EmployeeOfTheMonth *Employee = &Bob
	EmployeeOfTheMonth.Address = "New Address 2"

	// OR

	(*EmployeeOfTheMonth).Address = "New Address 3"

	fmt.Println(Bob)

}
