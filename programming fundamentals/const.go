package main

import "fmt"

const (
	a = 42
	b = 43.52
	c = "Chris"
)

func constant() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
