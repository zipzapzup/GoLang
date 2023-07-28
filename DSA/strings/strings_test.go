package main

import "testing"

func TestCheckStrings(t *testing.T) {
	if checkStrings("hello", "o") == false {
		t.Error("test failed")
	}
}

func TestCheckStringsv2(t *testing.T) {
	// create test table
	var tests = []struct {
		input1 string
		input2 string
		result bool
	}{
		{
			input1: "test",
			input2: "e",
			result: true,
		},
		{
			input1: "ball",
			input2: "a",
			result: true,
		},
	}

	for _, testcases := range tests {
		if checkStrings(testcases.input1, testcases.input2) != testcases.result {
			t.Errorf("test is failing input1: %v, input2: %v is not equal to result: %v", testcases.input1, testcases.input2, testcases.result)
		}
	}

}
