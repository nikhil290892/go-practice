package main

import "fmt"

func printascii() { //print ascii values of the numbers
	for p := 33; p <= 122; p++ {
		fmt.Printf("%v\t\t%#U\n", p, p)
	}
}
