package main

import (
	"log"
)

/**
* https://twitter.com/psankar/status/999198622100045824
* ? is the output
*
 */
func main() {

	a1 := [3]int{1, 2, 3}
	a2 := [3]int{4, 5, 6}

	m := make(map[[3]int]bool)
	m[a1] = true
	m[a2] = true

	var b [][]int
	for i, _ := range m {
		log.Println("Adding: ", i[:])
		b = append(b, i[:])
	}
	log.Println("Final slice is: ", b)
}
