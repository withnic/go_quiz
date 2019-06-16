package main

import "fmt"

/*
https://twitter.com/silvr/status/1123668404932681730
Here's my crack at a #golang pop quiz (inspired by @davecheney
and others). What is the output of this program?

value is +50
value is 50
won't compile
value is 2

*/
func main() {
	value := 50
	fmt.Println(fmt.Sprintf("value is %s", string(value)))
}
