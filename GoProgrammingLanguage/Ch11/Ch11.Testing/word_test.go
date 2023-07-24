package word

import "testing"

// test for isPalindromev2
func TestIsPalindromev2(t *testing.T) {

	if isPalindromev2("kayak") == false {
		t.Error(`isPalindromev2("kayak") = false`)
	}

}

func TestNonisPalindromev2(t *testing.T) {
	if isPalindromev2("palindrome") == true {
		t.Error(`isPalindromev2("palindrome") = true`)
	}
}

// extend test package to test French
func TestFrenchPalindrome(t *testing.T) {
	if isPalindromev2("été") == false {
		t.Error(`isPalindromev2("été") == false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if isPalindromev2(input) == true {
		t.Errorf(`ispalindromev2( %q ) == false`, input)
	}
}

// more comprehensive test

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", false},
		{"Evil I did dwell; lewd did I live.", false},
		{"Able was I ere I saw Elba", false},
		{"été", true},
		{"Et se resservir, ivresse reste.", false},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome

	}

	for _, test := range tests {
		if got := isPalindromev2(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}

}
