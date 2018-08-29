package main

import "fmt"

//#golang quiz
// https://twitter.com/bot_golang/status/1024317026603421697
//What does the following program print?
// Complication error
// not equal
// equal
func main() {
	a := [...]int{5,6,7}
	b := [3]int{5,6,7}
	if a == b {
		fmt.Println("equal")
		return
	}
	fmt.Println("not equal")
}
