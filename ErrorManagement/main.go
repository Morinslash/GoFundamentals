package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("This error wraps the first one: %w", err)
	fmt.Println(err2)
}
