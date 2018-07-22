package main

import "fmt"

/**
https://twitter.com/kchristidis_/status/1013242397566922752
What does this print?
panic: index out of range
[]
*/
func main() {
	a := []int{0, 1}
	fmt.Printf("%v", a[len(a):])
}
