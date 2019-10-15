package main

import "fmt"

func ex4() {
	bd := 1992
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
