package main

import "fmt"

func slicerange() {
	x := []int{4, 34, 67534, 64, 534}
	fmt.Println(x)
	//get the array value by range
	for i, v := range x {
		fmt.Println(i, v)
	}
	//get the array value by index
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
}
