package main

import "fmt"

func myFunction(firstName string, lastName string) string {
	fullName := firstName + " " + lastName
	return fullName
}

func main() {
	fmt.Println("Hello world")

	fullName := myFunction("Jethro", "May")
	fmt.Println(fullName)
}
