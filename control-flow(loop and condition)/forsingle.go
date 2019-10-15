package main

import "fmt"

func forst() {
	a := 2
	b := 3
	for a < b { //for single statement with single condition
		a = a * 2
	}
	fmt.Println(a)
}
func fornot() { //This is pretty much like the for in C language but not same and implements while within it
	c := 1
	for {
		if c > 9 {
			break
		}
		fmt.Println(c)
		c++
	}
	fmt.Println("done.")
}
