package main

import (
	"fmt"
	"time"
)

// https://twitter.com/davecheney/status/1083215038939258880
// #golang pop quiz. What does this program print?
//0.25s
//0s
//250ms
//nothing, doesn't compile
func main() {
	fmt.Println(0.25 * time.Second)
}
