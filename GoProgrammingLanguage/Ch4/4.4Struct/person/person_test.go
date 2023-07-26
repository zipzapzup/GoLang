package main

import (
	"fmt"
	"testing"
)

func TestCountPeople(t *testing.T) {

	fmt.Println("Starting Test")
	newHouse := new(House)
	p1 := person{"John", 20}
	p2 := person{"Lisa", 20}
	p3 := person{"Bart", 20}
	newHouse.storeHouse("House of Dior", p1)
	newHouse.storeHouse("House of Dior", p2)
	newHouse.storeHouse("House of Dior", p3)

	want := 3
	got := newHouse.countPeople()

	if got != want {
		t.Fatalf("newHouse.countPeople() does not match got: %q, want: %q", want, got)
	}

}

func TestAddNumber(t *testing.T) {
	// write test cases one by one
	if addNumber(2, 3) != 5 {
		t.Errorf(`addNumber(2, 3) != 5 is false`)
	}
	if addNumber(10, 10) != 20 {
		t.Errorf(`addNumber(10,10) != 20 is false`)
	}
	if addNumber(0, 10) != 10 {
		t.Errorf(`addNumber(0, 10) != 10 is false`)
	}
	if addNumber(-5, 10) != 5 {
		t.Errorf(`addNumber(-5, 10) != 5 is false`)
	}

	// write comprehensive set of cases
	type test []struct {
		input1 int
		input2 int
		result int
	}

	tabletest := test{
		{
			input1: 2,
			input2: 2,
			result: 4,
		},
		{
			input1: 3,
			input2: 4,
			result: 7,
		},
	}

	for _, test := range tabletest {
		if addNumber(test.input1, test.input2) != test.result {
			t.Errorf("addNumber(%q, %q) is not equal to %q", test.input1, test.input2, test.result)
		}
	}

}
