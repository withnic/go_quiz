package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/959196014480474114
#golang pop quiz, what does this program print?
* 12, 16
* 12, 12
* depends on GOARCH
* doesn't compile
*/
func main() {
	s := "It's Friday!"
	fmt.Println(len(s), cap(s))
}
