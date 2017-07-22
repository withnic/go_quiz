package main

import "fmt"

/*
https://twitter.com/nickdsnyder/status/875830289850748928
* prints whats?
* [a b c]
* [a b d]
* [a b c d]
* panic
*/
func main() {
	x := make([]string, 0, 3)
	x = append(x, "a", "b")
	y := append(x, "c")
	_ = append(x, "d")
	fmt.Print(y) // ?
}
