package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/989204013269835781
* What does this golang program print?
* "ab"
* "aa"
* "bb"
* "ba"
 */
func main() {
	s := []byte("")

	s1 := append(s, 'a')
	s2 := append(s, 'b')

	fmt.Print(string(s1), string(s2))
}
