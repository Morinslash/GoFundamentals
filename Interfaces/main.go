package main

import (
	"bytes"
	"fmt"
	"strings"
)

type printer interface {
	Print() string
}

type user struct {
	username string
	id       int
}

type menuItem struct {
	name   string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}
	return b.String()
}

func (u user) Print() string {
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

func main() {
	var p printer
	p = user{username: "Adam", id: 1}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee", prices: map[string]float64{
		"small":  1.65,
		"medium": 1.80,
		"large":  19.50,
	},
	}
	fmt.Println(p.Print())

	//u := p.(user) // =>this will cause Panic

	u, ok := p.(user)
	fmt.Println(u, ok)
	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	switch v := p.(type) {
	case user:
		fmt.Println("Found a user!", v)
	case menuItem:
		fmt.Println("Found menu item!", v)
	default:
		fmt.Println("Not sure what this is...")
	}
}
