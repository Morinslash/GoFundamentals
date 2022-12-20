package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1) // register async task

	go func() {
		fmt.Println("This happens asynchronously")
		waitGroup.Done() // confirm execution
	}()

	fmt.Println("This is synchronous")
	waitGroup.Wait() // wait till all async done
}
