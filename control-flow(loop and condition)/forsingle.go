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
func fornot() {
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
