package main

import "fmt"

func addOne() func() int {
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
		// this anonymous function
		// has access to the x variable
		// defined in the parent function
		x++
		return x + 1
	}
}

func main() {
	myFunc := addOne()
	fmt.Println(myFunc())
}
