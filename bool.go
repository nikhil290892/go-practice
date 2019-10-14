package main

import (
	"fmt"
)

var f bool
var no int

func booleancheck() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			f = true
			fmt.Println(f)
			no = no + 1
		}
	}
	fmt.Println(no)

}
