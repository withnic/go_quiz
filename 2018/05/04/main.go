package main

import (
	"fmt"
	"unsafe"
)

/**
* https://twitter.com/inancgumus/status/992071880982122497
* What's the total size of this struct?
* // tip: is 1byte and int64 is 8bytes
*
struct {
	a int64
	b bool
	c bool
}
*
* 10 bytes
* 16 bytes
* 32 bytes
* Depends
*/
func main() {
	var v struct {
		a int64
		b bool
		c bool
	}
	fmt.Print(unsafe.Sizeof(v))
}
