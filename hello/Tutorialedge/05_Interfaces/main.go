package main

import "fmt"

func myFunc(a interface{}) {
	fmt.Println(a)
}

type guitarist interface {
	PlayGuitar()
}

type baseGuitarist struct {
	name string
}

type acousticGuitarist struct {
	name string
}

func (b baseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Base Guitar\n", b.name)
}

func (b acousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the acoustic Guitar\n", b.name)
}

func main() {
	var myName string
	myName = "Véronique"
	myFunc(myName)

	var player baseGuitarist
	var player2 acousticGuitarist
	player.name = "Bob"
	player2.name = "Rocky"
	player.PlayGuitar()

	var guitarists []guitarist // array of type "guitarist"
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)
	fmt.Println(guitarists)

	for i := 0; i < len(guitarists); i++ {
		fmt.Println(guitarists[i])
	}

	// taking an array
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("The elements of the array are: ")

	// using for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
