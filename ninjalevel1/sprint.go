package main

import (
	"fmt"
)

var p int = 42
var q string = "James Bond"
var r bool = true

func stringprint() {
	s := fmt.Sprintf("%v%v%v", p, q, r)
	fmt.Println(s)
}
