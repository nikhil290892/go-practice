package main

import "fmt"

func slicemultidim() {
	jb := []string{"James", "Bond"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
