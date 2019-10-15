package main

import "fmt"

var x int

func ex5() {
	for i := 10; i <= 100; i++ {
		x = i % 4
		fmt.Println("When", i, "is divided by 4, the modulus is\n", x)
	}
}
