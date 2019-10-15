package main

import "fmt"

func bnc() { //break and continue example
	d := 0
	for {
		d++
		if d > 100 {
			break
		}
		if d%2 != 0 {
			continue
		}
		fmt.Println(d)
	}
	fmt.Println("Even numbers printed.")
}
