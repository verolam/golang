package main

import "fmt"

func myfunction(firstName string, lastName string) (string, error) {
	fullName := firstName + " " + lastName
	return fullName, nil
}

func main() {
	fmt.Println("Hello World!! ")
	fullName, err := myfunction("Véronique", "Lamrani")
	if err != nil {
		fmt.Println("HAndle Error Case")
	}

	fmt.Println("My name is " + fullName)
}
