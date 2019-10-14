package main

import "fmt"

func expression() {
	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 42)
	d := (41 < 42)
	e := (42 > 41)
	f := (42 != 43)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
