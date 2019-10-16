package main

import "fmt"

func slicedel() {
	x := []int{4, 63, 3, 43, 65}
	fmt.Println(x)
	x = append(x, 32, 76, 56, 45, 65, 74)
	fmt.Println(x)
	y := []int{52352, 7373, 623652, 6262}
	x = append(x, y...)
	fmt.Println(x)
	x = append(x[:1], x[4:]...)
	fmt.Println(x)
}
