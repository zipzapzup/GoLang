package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {

	fmt.Println("Tree Struct")

	var x tree
	x.value = 10
	// establish root

	// this adds to the left of the tree (Value: 10, L: 5, R: NIL)
	add(5, &x)
	add(99, &x)
	add(20, &x)

	strings.Split()

	fmt.Println(x)

	fmt.Println(appendValue([]int{}, &x))

}

// add version 1
// no need to return since the pointer will modify the value
// modifying the value indirectly via pointer
func add(v int, t *tree) {

	if t == nil {
		t = new(tree)
		t.value = v
	}
	newtree := new(tree)
	newtree.value = v
	if v < t.value {
		t.left = newtree

	} else {
		t.right = newtree
	}
}

// func to add the value to a slice
func appendValue(v []int, t *tree) []int {

	if t != nil {

		v = appendValue(v, t.left)
		v = append(v, t.value)
		v = appendValue(v, t.right)

	}
	return v
}

// func printTree(t *tree) {
// 	var newSlice []int

// 	if t == nil {
// 		return
// 	}
// }
