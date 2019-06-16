package main

import (
	"fmt"
	"math"
)

// https://twitter.com/hajimehoshi/status/1129679065500139522
func main() {
	a := map[float64]int{}
	a[math.Acos(8)] = 1
	a[math.Acos(8)] = 2
	a[math.Acos(8)] = 3
	i := 0
	for range a {
		i++
	}
	fmt.Println(i)
	fmt.Println(len(a))
}
