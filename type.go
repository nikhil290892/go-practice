package main

import (
	"fmt"
)

type hotdog int //declare my own type

var a int
var b hotdog

func vartype() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b) //convert type hotdog to type int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
