package main

import (
	"fmt"
)

// https://twitter.com/bot_golang/status/1131236764083838976
// #golang quiz
// What does the following program print?
// equal
// not equal
// compilation error
func main() {
	a := []string{"a", "b"}
	b := []string{"a", "b"}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
