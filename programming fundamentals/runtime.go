package main

import (
	"fmt"
	"runtime"
)

func run() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
