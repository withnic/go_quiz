package main

import (
	"fmt"
)

// https://twitter.com/davecheney/status/1068347813531070464
// #golang pop quiz: what does this program print?
//1.000000001e+10
//(1e+10+10i)
//nada, doesn't compile
func main() {
	x := 1e10 + 1e1i
	fmt.Println(x)
}
