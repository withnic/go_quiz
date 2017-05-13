package main

import "fmt"

// @davecheney Pop quiz, what does this print, and why ?
// https://twitter.com/davecheney/status/364215899340804096
func main() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}
