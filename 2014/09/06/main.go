package main

import "fmt"

/**
Pop quiz, what does this #golang program print (no cheating)
https://twitter.com/davecheney/status/508076305003196416
**/
func main() {
	a := [...]int{5, 4: 1, 0, 2: 3, 2, 1: 4}
	fmt.Println(a)
}
