package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/959233475193667584
Friday extra #golang pop quiz, what does this program print?
* 0
* doesn’t compile
* 10
* 11
*/
func main() {
	i := 0
	for i := 0; i <= 10; i++ {
	}
	fmt.Println(i)
}
