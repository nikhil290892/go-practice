package main

import "fmt"

func sliceappend() {
	x := []int{4, 63, 3, 43, 65}
	fmt.Println(x)
	x = append(x, 32, 76, 56, 45, 65, 74)
	fmt.Println(x)
	y := []int{52352, 7373, 623652, 6262}
	x = append(x, y...) //...after any variable array type appends all the value of that array to the declared variable
	fmt.Println(x)
}
