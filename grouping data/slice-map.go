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
	//to add a new value to the map
	m["Phil"] = 38
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m)
	//delete an entry in map
	delete(m, "James")
	fmt.Println(m)
}
