package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	myData := []byte("All the data I wish to add to the file")

	// WriteFile returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.data", myData, 0777)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
