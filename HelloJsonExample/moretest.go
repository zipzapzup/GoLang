package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	JsonContent := `{
			"2023-06-30": {
				"1. open": "134.69",
				"2. high": "135.03",
				"3. low": "133.425",
				"4. close": "133.81",
				"5. adjusted close": "133.81",
				"6. volume": "4236677",
				"7. dividend amount": "0.0000",
				"8. split coefficient": "1.0"
			},
			"2023-06-29": {
				"1. open": "131.75",
				"2. high": "134.35",
				"3. low": "131.69",
				"4. close": "134.06",
				"5. adjusted close": "134.06",
				"6. volume": "3639836",
				"7. dividend amount": "0.0000",
				"8. split coefficient": "1.0"
			},
			"2023-06-28": {
				"1. open": "132.06",
				"2. high": "132.17",
				"3. low": "130.91",
				"4. close": "131.76",
				"5. adjusted close": "131.76",
				"6. volume": "2753779",
				"7. dividend amount": "0.0000",
				"8. split coefficient": "1.0"
			}
	}
	`

	fmt.Println("Hello, 世界\n", JsonContent)

	var f interface{}
	err := json.Unmarshal([]byte(JsonContent), &f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	fmt.Println("Part 2\n \n")

	assertf := f.(map[string]interface{})
	for k, v := range assertf

}
