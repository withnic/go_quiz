package main

import "fmt"

// @mhemmings 2017/04/15
func main() {
	// What 3 integers does this print?
	fmt.Println(foo())
}

func foo() (int, int, int) {
	var x int
	return x + x, func() int {
			x += 1
			defer func() {
				x += 10
			}()
			return x
		}(),
		func() int {
			x += 2
			defer func() {
				x += 20
			}()
			return x
		}()
}
