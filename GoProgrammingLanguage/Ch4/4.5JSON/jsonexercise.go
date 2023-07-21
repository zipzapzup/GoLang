package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	fmt.Println("JSON Hello, World")

	// JSON can convert Go Object to JSON and this is known as Marshalling

	type Employer struct {
		Name      string
		Founded   int `json:"founded"`
		Employees []string
	}

	var IBM Employer = Employer{
		Name: "IBM", Founded: 1900, Employees: []string{"John Doe", "Alex AO", "Nameless Tree"},
	}
	fmt.Println(IBM)

	// use Marshal
	result, err := json.Marshal(IBM)
	if err != nil {
		log.Fatalf("JSON Marshalling Failed %s", err)
	}
	fmt.Printf("\nJSON Output:\n%s", result)

	// better format use Marshal Indent
	// Marshal Indent can format however it needs to be in string
	// json is in byte and converted to unintrepreted %s

	data := map[string]int{
		"ABC": 20,
		"CED": 10,
	}
	b, error := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("\nJSON Output\n%s", error)
	}
	fmt.Println("\n", string(b))

	b, error = json.MarshalIndent(IBM, "", " ")
	if err != nil {
		log.Fatalf("\nJSON ERROR: %s", err)
	}
	fmt.Println("UNMARSHAL with Indent:\n", string(b))
}
