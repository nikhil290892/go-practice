package main

import "fmt"

func nestedloop() {
	for a := 0; a < 10; a++ {
		for b := 1; b < 5; b++ {
			fmt.Println("Nikhil is great", "inner loop", b, "outer loop", a)
		}
	}
}
