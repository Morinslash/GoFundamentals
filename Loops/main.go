package main

import "fmt"

func main() {
	//InfiniteLoop()
	//ConditionLoop()
	//CounterLoop()
	CollectionLoop()
}

func CollectionLoop() {
	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Done!")
}

func CounterLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func ConditionLoop() {
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func InfiniteLoop() {
	// infinity loop
	i := 99
	for {
		fmt.Println(i)
		i += 1
		break
	}
}
