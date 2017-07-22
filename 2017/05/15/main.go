package main

import "errors"

type e struct{}

func (*e) Error() string { return "boom" }

func f() *e { return nil }

/**
#golang pop quiz: what does this program print?
true
false
whoops
boom
https://twitter.com/davecheney/status/862793327321665536
*/
func main() {
	err := errors.New("whoops")
	err = f()
	println(err == nil)
}
