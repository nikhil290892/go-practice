package main

import "fmt"

const (
	e = iota
	h
	g
)
const (
	j = iota
	k
	l
)

func iotacheck() {
	fmt.Println(e)
	fmt.Println(h)
	fmt.Println(g)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

//successive const integer is iota. The constants are assigned a successive value based on their declaration
