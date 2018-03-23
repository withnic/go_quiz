package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/957545937664065536
#golang pop quiz: what does this program print?
* true
* false
* implementation dependant
*/
func main() {
	var a bool
	defer fmt.Println(a == true)
	a = !a
}
