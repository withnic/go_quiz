package main

import "fmt"

/**
@wadey
https://twitter.com/wadey/status/385475359350603776
Golang quiz time: Check your understand of scope and values to figure out what is going on here:
**/

type A struct {
	Foo string
}

func main() {
	as := []A{
		{"1"},
		{"2"},
		{"3"},
	}

	// Create []*A from []A
	as2 := make([]*A, len(as))
	for i, a := range as {
		as2[i] = &a
	}

	fmt.Printf("%v %v %v\n", as[0], as[1], as[2])
	fmt.Printf("%v %v %v\n", as2[0], as2[1], as2[2])
}
