package main

import "fmt"

// https://twitter.com/fzerorubigd/status/982850150157705216
// result of this
/**
* a
* b
* c
 */
func main() {
	var a = []string{"a", "b", "c"}
	var b = make([]*string, len(a))
	for i, j := range a {
		b[i] = &j
	}
	fmt.Println(*b[1])
}
