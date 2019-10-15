package main

import "fmt"

func condition() {
	if true {
		fmt.Println("Hello Nikhil")
	}
	if false {
		fmt.Println("Byee Nikhil")
	}
	if !true {
		fmt.Println("Byee Nikhil")
	}
	if !false {
		fmt.Println("Hello Nikhil")
	}
}
func conditioninit() {
	if n := 42; n == 2 {
		fmt.Println("Howdy")
	}
	fmt.Println("The initialized and compared value are different")
}
