package main

import (
	"fmt"
)

var whatIsThe = answerToLife()

func answerToLife() int {
	return 50
}

func init() {
	whatIsThe = 0
}

func main() {
	if whatIsThe == 0 {
		fmt.Println("Its all a lie")
	}
}
