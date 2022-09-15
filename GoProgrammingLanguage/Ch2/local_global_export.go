package main

import "fmt"

// golang has 3 concepts of variable
// local
// global
// exported

var globalVar string = "global variable"

// any variable declared outside the function is global

var ExportVar string = "exported variable"

// variable is exported due to the capital case of E

func main() {

	// Golang has the concept of local variable
	// any variable declared inside the function is local

	var localVar string = "local variable"
	fmt.Println(globalVar, "&", localVar)
}
