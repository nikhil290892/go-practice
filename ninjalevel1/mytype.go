package main

import (
	"fmt"
)

type random int

var t random

func mytype() {
	fmt.Println(t)
	fmt.Printf("%T\n", t)
	t = 42
	fmt.Println(t)
}
