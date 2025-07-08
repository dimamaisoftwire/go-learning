package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	fmt.Println(greeting(name))
}

func arrayContains(list []string, element string) bool {
	// iterate over the array and compare given string to each element
	for _, value := range list {
		if value == element {
			return true
		}
	}
	return false
}

func greeting(name string) string {
	formattedName := fmt.Sprintf("%.20s", name) // Format so that 20 characters max

	var specialNames = []string{} // Special names that require thanks
	specialNames = append(specialNames, "Rob Pike")
	specialNames = append(specialNames, "Ken Thompson")
	specialNames = append(specialNames, "Robert Griesemer")

	firstName := strings.Split(formattedName, " ")[0] // Get first word

	output := "Hello, " + firstName // Standard hello

	if arrayContains(specialNames, name) {
		output = output + ". Thanks for creating me!"
	}

	return output
}
