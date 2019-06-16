package main

import "fmt"

type integer int

const (
	one integer = iota
	two
)

// https://twitter.com/codingsince1985/status/1124575623941255169
// #golang quiz. what's the output?
// 2
// two
// error: index out of range
func main() {
	var i integer = 2
	fmt.Println(i)
}
