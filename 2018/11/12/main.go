package main

import "fmt"

type number struct{ value int }

//go:noinline
func (n *number) Add(x int) { n.value += x }

// https://twitter.com/quasilyte/status/1062000322153975811
// #golang quiz
//What does the code below prints?
// 10
// 15
// 5
func main() {
	n := &number{value: 10}
	(*n).Add(5)          // Note the explicit dereference
	fmt.Println(n.value) // 15 or 10?
}
