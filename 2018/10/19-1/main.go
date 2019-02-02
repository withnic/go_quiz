package main

import (
"fmt"
)

// https://twitter.com/davecheney/status/1053419185680744448
// #golang pop quiz: what does this program print?
// true
// false
// panic
func main() {
var a []int
b := a[:]
fmt.Println(b == nil)
}
