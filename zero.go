package main

import (
	"fmt"
)

var s string
var n int

//DECLARE a variable of certain type
//and then assign the value of that type to the variable

func zerotype() {
	fmt.Println("Printing the value of y", s)
	fmt.Printf("%T\n", s)
	s = "Bond, James Bond"

	fmt.Println(n)
	fmt.Printf("%T\n", n)
}
