// Function to fetch the URL based on command arguments
// go run fetch_original.go https://google.com https://yahoo.com

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("No URLs Supplied")
		os.Exit(1)
	}

	fmt.Println("Starting Loop")
	for _, url := range os.Args[1:] {
		fmt.Println("\n\nFetching the following: ", url)
		fetch(url)

	}

	fmt.Println("\nFinish of Execution")

}

func fetch(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic("HTTP Get Request failed")
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer resp.Body.Close()
	secs := time.Since(start).Seconds()
	fmt.Println("Duration", secs, "Seconds")
}
