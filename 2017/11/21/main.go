package main

import "fmt"

/**
#golang pop quiz: what does this program print?
https://twitter.com/davecheney/status/932662557810221057
*/
func main() {
	x := []int{1, 2, 3}
	var i int
	for i = range x {
		x = nil
	}
	fmt.Println(i)
}
