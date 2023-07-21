package main

import "fmt"

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

	fmt.Println(x)

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

// func printTree(t *tree) {
// 	var newSlice []int

// 	if t == nil {
// 		return
// 	}
// }
