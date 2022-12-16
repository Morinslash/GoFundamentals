package main

import "fmt"

func main() {
	//testScores := []float64{
	//	87.3,
	//	105,
	//	63.5,
	//	27,
	//}

	testScores := map[string]float64{
		"Harry":    87.3,
		"Hermione": 105,
		"Ronald":   63.5,
		"Neville":  27,
	}

	c := clone(testScores)

	//fmt.Println(&testScores[0], &c[0], c)
	fmt.Println(c)
}

// Simple generics
//func clone[V any](s []V) []V { // => here we are using generics for the clone
//	result := make([]V, len(s))
//	for i, v := range s {
//		result[i] = v
//	}
//	return result
//}

// Maps Generic
func clone[K comparable, V any](m map[K]V) map[K]V { // => keys in map MUST be comparable
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}
