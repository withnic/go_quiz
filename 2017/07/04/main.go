package main

import "log"

/*
⎛⚆_⚆⎠ Why this works?
ʕ⚆ϖ⚆ʔ I dunno!
https://twitter.com/__dc0d__/status/882217423356690433
*/
func main() {
	a := []int{1}
	x, a := a[0], a[1:]
	log.Println(x, a)

}
