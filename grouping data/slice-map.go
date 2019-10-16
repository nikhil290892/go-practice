package main

import "fmt"

func slicemap() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 29,
		"Nikhil":     27,
	} //composite literal
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Shawn"])
	v, ok := m["Shawn"] //check if the key exists as this index in the map
	fmt.Println(v)
	fmt.Println(ok)
}
