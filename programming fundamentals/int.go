package main

import (
	"fmt"
)

var u int
var v float64

func intcheck() {
	u = 45
	v = 43.653652
	fmt.Println(u)
	fmt.Println(v)
	fmt.Printf("%T\n", u)
	fmt.Printf("%T\n", v)
}

//byte is an alias for uint8
//rune is an alias for int32
//GO uses UTF-8 encoding mechanism
