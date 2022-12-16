package main

import "fmt"

func main() {
	//SimpleIf()
	//Else()
	//ElseIf()
	//IfElseWithInitializer()
	//BasicSwitch()
	LogicalSwitch()
}

func LogicalSwitch() {
	//switch i := 8; true {
	switch i := 8; { // <= true implicit
	case i < 5:
		fmt.Println("i is less than 5")
	case i < 10:
		fmt.Println("i is less than 10")
	default:
		fmt.Println("is is greater than or queal to 10")
	}
}

func BasicSwitch() {
	//i := 99
	switch i := 99; i {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*i + 3:
		fmt.Println("second case")
	default:
		fmt.Println("default case")
	}
}

func IfElseWithInitializer() {
	if i := 15; i < 5 {
		fmt.Println("is is less than 5")
	} else if i < 10 {
		fmt.Println("i less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("after the if statement")
}

func ElseIf() {
	i := 5
	if i < 5 {
		fmt.Println("is is less than 5")
	} else if i < 10 {
		fmt.Println("i less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("after the if statement")
}

func Else() {
	i := 4
	if i < 5 {
		fmt.Println("i is less than 5")
	} else {
		fmt.Println("i is at least 5")
	}
	fmt.Println("after the if statement")
}

func SimpleIf() {
	i := 4
	if i < 5 {
		fmt.Println("i is less than 5")
	}
	fmt.Println("after the if statement")
}
