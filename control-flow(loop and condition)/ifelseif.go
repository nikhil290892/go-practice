package main

import "fmt"

func ifelseifelse() {
	x := 42
	if x == 42 {
		fmt.Println(x)
	} else if x == 32 {
		fmt.Println("Sorry")
	} else {
		fmt.Println("You're great")
	}
}
