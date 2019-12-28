package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func stru() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   16,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}
