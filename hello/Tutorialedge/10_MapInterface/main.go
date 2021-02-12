package main

import "fmt"

type Service interface {
	sayHi()
}

type MyService struct{}

func (s MyService) sayHi() {
	fmt.Println("Hi")
}

type SecondService struct{}

func (s SecondService) sayHi() {
	fmt.Println("Hello 2nd servcie")
}

func main() {
	fmt.Println("Go Maps Tutorial")

	interfaceMap := make(map[string]Service)

	// We can populate our map with simple ids to particular service
	interfaceMap["SERVICE-ID-1"] = MyService{}
	interfaceMap["SERVICE-ID-2"] = SecondService{}

	// Incoming HTTP Request wants service 2
	// We can use the incoming uuid to lookup the required service
	// and call its SayHi() method

	interfaceMap["SERVICE-ID-1"].sayHi()

	// Iteration thru interfaceMap
	for key, service := range interfaceMap {
		fmt.Println(key)
		service.sayHi()
	}

}
