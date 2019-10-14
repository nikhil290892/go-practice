package main

import "fmt"

func str() {
	s := "Hello, Playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	s2 := `Hey, Nikhil. How are you doing?
	I heard you are excelling in GO` //String literals will keep the string as it is and not remove the whitespace
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) //Print the UTF-8 character representer
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("At index position %d, we have hex %#x\n", i, v)
	}
}
