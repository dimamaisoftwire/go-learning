package main

import (
	"bufio"
	"fmt"
	"os"
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

func greeting(name string) string {
	fname := fmt.Sprintf("%.20s", name)
	output := "Hello, " + fname
	if name == "Robert Griesemer" {
		output = output + ". Thanks for creating me!"
	}
	return output
}
