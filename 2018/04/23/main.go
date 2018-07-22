package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/988087571975757830
* What does this golang program print and why ?
* 0
* 1
* 32
* panic
 */
func main() {
	s := ""
	r := []rune(s)
	fmt.Print(len(r[:32]))
}
