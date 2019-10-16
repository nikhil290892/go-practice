package main

import "fmt"

func slicemake() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[1] = 32
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 32141) //to increase the length of the array use append and the compiler knows
	//that this array has a capacity of 100 so it will not create a new array->copy the values->drop the old array.
	//It can use the same array with the help of slice and append.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 32141, 31234, 41423, 352462) // when the array capacity is reached, the compiler, during the runtime,
	//will create a new array->copy the values->drop the old array. And since it knows the capacity has been breached,
	//it will double the capacity.
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
