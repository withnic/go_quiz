package main

import (
	"fmt"
)

type precision string

/**
#golang pop quiz
https://twitter.com/webRat/status/951939708778631169
* Doesn't compile
* Panic
* Empty String
*/
func main() {
	fmt.Println(precision(1))
}

// cc @davecheney
