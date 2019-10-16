package main

import "fmt"

func arrayintro() {
	var x [5]int //initialize the array with the size
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
