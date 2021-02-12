package main

import (
	"fmt"
)

func main() {

	var s string = "Hello World \n"
	fmt.Println(s)

	for index, character := range s {
		fmt.Printf("The character %c is in posistion %d \n", character, index)
	}
}
