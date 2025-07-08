package main

import (
	"testing"
)

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Bob", "Hello, Bob"},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Mynameiswayyyyyyytoooo longg", "Hello, Mynameiswayyyyyyytoo"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"abcdcba", "Hello, abcdcba. Cool, a palindromic name!"},
	{"abccba", "Hello, abccba. Cool, a palindromic name!"},
	{"abccba fs", "Hello, abccba"},
	{"a", "Hello, a. Cool, a palindromic name!"},
}

var inputTests = []struct {
	input    string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"", "Ok, no greeting for you"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("%s", "incorrect greeting, for name "+test.name+" got: "+result+" expected: "+test.expected)
		}
	}
}
func TestHadleInput(t *testing.T) {
	for _, test := range inputTests {
		result := hadleInput(test.input)
		if result != test.expected {
			t.Errorf("%s", "incorrect greeting, for name "+test.input+" got: "+result+" expected: "+test.expected)
		}
	}
}
