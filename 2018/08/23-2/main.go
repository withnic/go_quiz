package main

import "fmt"

// https://twitter.com/dmitshur/status/1032631272722624513
// true true
// true false
// true panic
// Undefined behavior
func main() {
	var s []int
	fmt.Println(s == nil)
	s = s[:]
	fmt.Println(s == nil)
}