package main

import "testing"

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Bob", "Hello, Bob"},
	{"Robert Griesemer", "Hello, Robert Griesemer. Thanks for creating me!"},
	{"My name is wayyyyyyy toooo longg", "Hello, My name is wayyyyyyy"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("%s", "incorrect greeting, for name "+test.name+" got: "+result+" expected: "+test.expected)
		}
	}
}
