package main

import "fmt"

type agent struct {
	first string
	last  string
	age   int
}
type secretagent struct {
	agent
	ltk bool
}

func embedstru() {
	sa1 := secretagent{
		agent: agent{
			first: "James",
			last:  "Bond",
			age:   35,
		},
		ltk: true,
	}
	sa2 := secretagent{
		agent: agent{
			first: "Miss",
			last:  "Moneypenny",
			age:   32,
		},
		ltk: false,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa1.age, sa2.ltk)
}
