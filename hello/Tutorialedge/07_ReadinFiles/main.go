package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("localfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))

	// content we want to write in the new file "myfile.data"
	mydata := []byte("all my data I want to write to a file")

	err2 := ioutil.WriteFile("myfile.data", mydata, 0777)
	if err2 != nil {
		fmt.Println(err2)
	}

}
