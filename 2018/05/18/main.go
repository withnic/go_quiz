package main

import (
	"fmt"
)

const (
	a = 42
	b = iota
	c = iota
)

/**
https://twitter.com/the_sttts/status/997460732349440000
* What's print?
*
*/
func main() {
	fmt.Println(a, b, c)
}
