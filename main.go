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

func hadleInput(input string) string {
	var output string
	if input != "" {
		output = greeting(input)
	} else {
		output = "Ok, no greeting for you"
	}
	return output
}

func main() {
	name := getName()
	message := hadleInput(name)
	fmt.Println(message)
}

func arrayContains(list []string, element string) bool {
	for _, value := range list {
		if value == element {
			return true
		}
	}
	return false
}

func isPalindrome(name string) bool {
	length := len(name)
	for i := 0; i < length/2; i++ {
		if name[i] != name[length-i-1] {
			return false
		}
	}
	return true
}

func greeting(name string) string {
	formattedName := fmt.Sprintf("%.20s", name)

	var specialNames = []string{}
	specialNames = append(specialNames, "Rob Pike")
	specialNames = append(specialNames, "Ken Thompson")
	specialNames = append(specialNames, "Robert Griesemer")

	firstName := strings.Split(formattedName, " ")[0]

	output := "Hello, " + firstName

	if arrayContains(specialNames, name) {
		output = output + ". Thanks for creating me!"
	}
	if isPalindrome(name) {
		output = output + ". Cool, a palindromic name!"
	}

	return output
}
