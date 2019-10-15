package main

import "fmt"

func ex7() {
	x := "Nikhil"
	if x == "Nikhil" {
		fmt.Println(true)
	} else if x == "Nik" {
		fmt.Println(false)
	} else {
		fmt.Println(false)
	}
}
