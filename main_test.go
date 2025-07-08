package main

import "testing"

func TestGreeting(t *testing.T) {
	var result, expected string
	result = greeting()
	expected = "Hello, world."
	if result != expected {
		t.Errorf("%s", "incorrect greeting, got: "+result+" expected: "+expected)
	}
}
