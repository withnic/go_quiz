package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/970441599913201669
* What does this program print? #golang #quiz (please don't use playground to answer)
* true
* false
* depends
 */
func main() {
	a, b := 1.0, 0.9

	const p, q = 1.0, 0.9

	fmt.Println((a - b) == (p - q))
}
