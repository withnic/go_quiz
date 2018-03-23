package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/956063763794485249
#golang pop quiz: what does this program print?
* 1
* 7
* 0
* 8
*/

func main() {
	const n = 7
	var a = []uint{n: n}
	fmt.Println(cap(a))
}
