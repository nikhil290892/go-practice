package main

import (
	"fmt"
)

//DECLARE there is a variable with the identifier "z"
//and that the variable with the identifier "z" is the "string"
//Assigns the "" of TYPE string to "z"
//false for booleans, 0 for integers, 0.0 for floats, "" for strings
//and nil for pointers, functions, interfaces, slices, channels and maps.

var y = 39
var z string

func keyword() {
	//short declaration operator
	//Declare a variable and assign a value of a certain type at the same time
	x := 54
	z = "Nikhil"
	fmt.Println("Value of x", x)
	fmt.Println("Value of global variable y", y)
	fmt.Println(z, "is the programmer for this")
}
