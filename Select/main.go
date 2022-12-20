package main

import "fmt"

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	for i := 0; i < 20; i++ {
		go func() {
			ch1 <- "message to channel 1"
		}()

		go func() {
			ch2 <- "message to channel 2"
		}()

		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
