package main

import "fmt"

func ex10() {
	//	fmt.Println(true && true)
	fmt.Println(true && false)
	//	fmt.Println(true || true)
	//	fmt.Println(false || false)
	fmt.Println(!true)
}

//./prog.go:8:36: redundant and: true && true
//./prog.go:10:36: redundant or: true || true
//Go vet exited.
