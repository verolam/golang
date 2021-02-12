package main

import (
	"fmt"
)

func main() {
	// Looping over arrays in Go

	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	// Loop in days , we ignore index returned by using an "_"
	for _, day := range days {
		fmt.Println(day)
	}
}
