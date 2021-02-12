package main

import "fmt"

// Basix Types
var myint int8 = 25
var myint2 int16 = 256

var myfloat1 float32 = 1135.255
var myfloat2 float32 = 23568.366

var mybool bool = true

var myName string = "Véronique"

const Pi = 3.14116

func main() {
	fmt.Println(myint, myint2)
	fmt.Println(myfloat1, myfloat2)
	fmt.Println(myfloat1 + 10)
	fmt.Println(myfloat1 == myfloat1+10) // it returns false
	fmt.Println(mybool)
	fmt.Println(myName)

}
