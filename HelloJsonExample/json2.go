package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	fmt.Println("HELLO")

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Could not marshal json: %s", err)
		return
	}

	fmt.Printf("json data: %s\n", jsonData)
}
