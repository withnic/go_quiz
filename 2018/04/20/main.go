package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/987344932192079872
* What dose this program print? You'll be surprised.
* true
* false
* panic
 */
func main() {
	x := []byte("a")
	func(a ...interface{}) {}(x)

	fmt.Print(cap(x) == cap([]byte("a")))
}
