package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := new(person)
	*p = person{"John", 20}

	p.changeName("Lucy")
	fmt.Println(p)
}

func (p *person) changeName(n string) {
	p.name = "NewName"
}
