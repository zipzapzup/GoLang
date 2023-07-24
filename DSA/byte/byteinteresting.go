package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Package Byte")

	var s string = "hello, world"
	fmt.Println(convertStringtoByte(s))

	fmt.Println(sortString("zcascab"))

}

func convertStringtoByte(s string) []byte {
	return []byte(s)
}

func sortString(s string) []string {
	list := strings.Split(s, "")
	sort.Strings(list)
	return list
}
