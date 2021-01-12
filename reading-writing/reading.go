package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read the contents of the file
	data, err := ioutil.ReadFile("localfile.data")

	// If it can't read it print out the error
	if err != nil {
		fmt.Println(err)
	}

	// If successful print out content of file as a string
	fmt.Println(string(data))
}
