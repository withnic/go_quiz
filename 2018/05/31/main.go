package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/1002170144079187968
* What this program prints?
* true
* false
* error: mismatched types
* panic
 */

func main() {
	x := interface {
		x()
		y()
	}(nil)

	y := interface {
		z()
		y()
		x()
	}(nil)
	fmt.Println(x == y)
}
