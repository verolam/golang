package main

import (
	"fmt"
	"sort"
)

func main() {
	//mymap := make(map[string]int)

	// Initialization of the map
	//	mymap["mykey"] = 10
	//mymap["zozo"] = 12

	mymap := map[string]int{
		"mykey": 25,
		"zozo":  26,
		"titi":  35,
	}

	fmt.Println(mymap["mykey"]) // Prints out 10

	// Iterating Keys and value
	for key, value := range mymap {
		fmt.Println(key)
		fmt.Println(value)
	}

	// Retrieving keys and create a new array keyArray
	var keyArray []string

	for key := range mymap {
		keyArray = append(keyArray, key)
	}
	fmt.Println(keyArray)

	// delete one element of map
	//delete(mymap, "mykey")
	//fmt.Println(mymap)

	// Checking if a key exists in a MAP
	if _, ok := mymap["zozo"]; ok {
		// The key "zozo" exixts in mymap
		fmt.Println(mymap["zozo"])

		if _, ok := mymap["zouzou"]; !ok {
			// The key "zozo" doesnt exist in mymap
			fmt.Println("key doesnt exist") // "key doesnt exist"
		}
	}

	// Sort Map by value by method sort

	sort.Strings(keyArray)

	for _, key := range keyArray {
		fmt.Println(key, mymap[key]) //elliot 35 	mykey 25 zozo 26
	}

}
