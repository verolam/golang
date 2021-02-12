package main

import "fmt"

//Employee is ...
type Employee struct {
	name string
}

// UpdateName method updates the e.name
func (e *Employee) UpdateName(newName string) {
	e.name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.name)
}

func main() {
	var employee1 Employee
	employee1.name = "Véronique"
	employee1.PrintName()
	employee1.UpdateName("Mouad")
	employee1.PrintName()
}
