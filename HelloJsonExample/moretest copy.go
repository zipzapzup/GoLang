package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type RawQuotes struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"6. volume"`
}

type rawDateSeries map[string]RawQuotes

type TimeSeries struct {
	DateSeries rawDateSeries `json:"Time Series (Daily)"`
}

func main() {

	JsonData := `{
		"Meta Data": {
			"1. Information": "Daily Time Series with Splits and Dividend Events",
			"2. Symbol": "IBM",
			"3. Last Refreshed": "2023-06-30",
			"4. Output Size": "Full size",
			"5. Time Zone": "US/Eastern"
		},
		"Time Series (Daily)": {
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
	}
	`

	fmt.Println("Hello, 世界\n", JsonData)
	Stock := TimeSeries{}

	json.Unmarshal([]byte(JsonData), &Stock)
	fmt.Println(Stock.DateSeries)
	fmt.Println(reflect.TypeOf(Stock.DateSeries["2023-06-28"]))
	a := Stock.DateSeries["2023-06-28"].Close
	fmt.Println(a)

}
