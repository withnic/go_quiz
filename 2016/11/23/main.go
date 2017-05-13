package main

import "fmt"

/**
#golang pop quiz: when iterating over a map, the order is
**/

// @davechenery https://twitter.com/davecheney/status/801174605998092289
func main() {

	// Slice
	fmt.Print("Slice: ")
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		print(v)
	}

	fmt.Print("\n-------\n")

	// Map
	fmt.Print("Map: ")
	m := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	for i := 1; i < 10; i++ {
		fmt.Printf("\n %d Times：", i)
		for _, v := range m {
			print(v)
		}
	}
}
