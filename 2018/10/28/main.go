package main

import "fmt"

// https://twitter.com/davecheney/status/1056487208989773824
// #golang pop quiz: what does this program print?
//ಠ_ಠ
//55
//7
//dude, that won’t compile
func main() {
	ಠ_ಠ := '7'
	fmt.Println(ಠ_ಠ)
}