package main

import "fmt"

// https://twitter.com/davecheney/status/1133288972875063297
// Too easy? What does this program print?
// a
// b
// c
// Something else(panic?)
func main() {
	var a, b *int
	var c = b
	switch c {
	case a:
		fmt.Println("a")
	case b:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
