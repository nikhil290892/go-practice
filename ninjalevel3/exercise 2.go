package main

import "fmt"

func ex2() {
	for i := 65; i <= 90; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%#U\t\n", i)
		}
	}
}
