package main

import "fmt"

type I interface{}
type J I
type X *struct{}

/**
https://twitter.com/the_sttts/status/948918766506860544
#golang pop quiz imitation: what does this program print?

*/
func main() {
	var x *struct{}
	fmt.Println(I(x) == J(x), I(x) == x, I(x) == nil, X(x) == nil)
}
