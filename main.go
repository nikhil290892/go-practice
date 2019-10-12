package main

import "fmt"

func main() {
	fmt.Println("Hello, Playground")
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
	operator()
	keyword()
}

func foo() {
	fmt.Println("I'm in foo.")
}

func bar() {
	fmt.Println("and then we exited")
}
