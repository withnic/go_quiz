package main

import "fmt"

/**
https://twitter.com/inancgumus/status/968650079954448385
What does this program print? #golang #quiz
* 0.01
* 0.0100000000000000000
* 0.0100000000000000002
* 0.0101000000000000000
*/
func main() {
	fmt.Printf("%.19f", 1.0-0.99)
}
