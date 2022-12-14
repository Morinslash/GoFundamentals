package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	//Arrays()
	//Slices()
	//Maps()
	Structs()
}

func Arrays() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(arr)

	fmt.Println(arr[1])
	arr[1] = "Chai Tea"
	fmt.Println(arr)

	arr2 := arr

	arr2[2] = "Chai Latte"

	fmt.Println(arr, arr2)
}
func Slices() {
	var s []string
	fmt.Println(s)

	s = []string{"Coffee", "Espresso", "Cappuccino"}
	fmt.Println(s)

	fmt.Println(s[1])
	s[1] = "Chai Tea"
	fmt.Println(s)

	fmt.Println("Append")
	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)
	fmt.Println("Delete")
	slices.Delete(s, 1, 2)
	fmt.Println(s)
	fmt.Println("Modify")
	s2 := s

	s2[2] = "Chai Latter"

	fmt.Println(s, s2)
}

func Maps() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latter"},
	}
	fmt.Println(m)
	fmt.Println(m["coffee"])
	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)
	// Delete
	delete(m, "tea")
	fmt.Println(m)
	//if removed key from the map, Go will still return empty map for this key
	fmt.Println(m["tea"])
	// to check if the key really exists use v,ok := query type and check ok value for true or false
	v, ok := m["tea"]
	fmt.Println(v, ok)

	// Reference type
	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)
}

func Structs() {

}
