package main

import "fmt"

// https://twitter.com/davecheney/status/1133172785440624640
// #golang pop quiz: what does this program print?
// a
// b
// c
// Something else
func main() {
	var a, b int
	var c = &b
	switch *c {
	case a:
		fmt.Println("a")
	case b:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
}
