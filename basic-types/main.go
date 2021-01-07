package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint)

}
