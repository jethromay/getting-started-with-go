package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Jethro"
	employee.UpdateName("May")
	employee.PrintName()
}
