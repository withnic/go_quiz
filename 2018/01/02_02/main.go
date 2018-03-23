package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/948154959316492288
#golang pop quiz: what does this program print?
1 true
2 false
3 uhhh, it doesn't compile
*/
func main() {
	a := -7 / 2
	b := -7 >> 1
	fmt.Println(a == b)
}
