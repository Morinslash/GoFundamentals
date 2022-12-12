package main

import "fmt"

func main() {
	s := "hello, world!"
	p := &s
	fmt.Println(p) // memory address
	fmt.Println(*p)

	*p = "hello, gophers!"
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p)

}
