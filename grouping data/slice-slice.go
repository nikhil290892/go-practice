package main

import "fmt"

func sliceslice() {
	x := []int{4, 63, 3, 43, 65}
	fmt.Println(x)
	fmt.Println(x[0:1])
	fmt.Println(x[1:])
	fmt.Println(x[2:4])
	for i := 0; i <= 4; i++ {
		if i != 2 {
			fmt.Println(x[i])
		}
	}
}
