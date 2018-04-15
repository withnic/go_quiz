package main

import (
	"fmt"
	"time"
)

/**
 * https://twitter.com/IcchaSethi/status/973652027153829889
 * Golang quiz: What is the output of this function?
 */
func main() {
	m, _ := time.ParseDuration("2500ms")
	x := time.Duration(m) * time.Millisecond
	fmt.Println(x)
}
