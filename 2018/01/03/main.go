package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/948335478402564096
#golang pop quiz: what does this program print?
* a
* -a
* 0
* nope, not gonna compile
*/
func main() {
	a := 'a'
	a += -a
	fmt.Println(a)
}
