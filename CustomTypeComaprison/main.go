package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Println("Sum of %v : %v\n", a1, s1)
	fmt.Println("Sum of %v : %v\n", a2, s2)
	fmt.Println("Sum of %v : %v\n", a3, s3)
}

type addable interface { //this is type interface, not behavioral
	int | float64 | string
}

func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
