package main

import "fmt"

// Agregate Types
// Arrays, Slices, Maps, Structs

func main() {
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Tuesday", "Friday", "Saturday", "Sunday"}
	weekdays := days[0:5] // SLice
	fmt.Println(days)
	fmt.Println(weekdays)

	// Maps
	mymap := map[string]int{
		"France": 68,
		"USA":    300,
		"Spain":  35,
	}

	fmt.Println(mymap["France"])
	for key, value := range mymap {
		fmt.Println(key, value)
	}

	// structs

	type Person struct {
		name string
		age  int
	}

	//	var myPerson Person
	myPerson1 := Person{"Véro", 50}
	myPerson2 := Person{"Mouad", 54}
	fmt.Println(myPerson1.name, myPerson2.age)

	// NEsted structs
	type Team struct {
		name    string
		players [2]Person
	}

	players := [...]Person{Person{name: "Forrest", age: 50}, Person{name: "NYC", age: 45}}
	fmt.Println(players)

	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)

}
