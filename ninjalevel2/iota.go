package main

import "fmt"

const (
	y1 = 2015 + iota
	y2 = 2015 + iota
	y3 = 2015 + iota
	y4 = 2015 + iota
)

func iotatest() {
	fmt.Println(y1, y2, y3, y4)
}
