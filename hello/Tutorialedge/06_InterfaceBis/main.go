package main

import "fmt"

type employee interface {
	language() string
}

type engineer struct {
	name string
	age  int
}

func (e engineer) language() string {
	return e.name + " programs in Go"
}

func main() {
	var engineer1 engineer
	engineer1.name = "Bob"
	fmt.Println(engineer1.language())
	var employees []employee
	employees = append(employees, engineer1)
	fmt.Println(employees)

}
