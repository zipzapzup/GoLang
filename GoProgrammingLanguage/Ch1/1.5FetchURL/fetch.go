// Fetch a URL and prints the content of the URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("URL not Found, Error:", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("HTTP Status Code: %s\n", resp.Status)
		fmt.Printf("%s", b)

	}
	secs := time.Since(start).Seconds()
	fmt.Printf("elapsed time: %.2fs \n", secs)

}
