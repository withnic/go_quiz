package main

import "fmt"

/**
* https://twitter.com/inancgumus/status/997114848986181637
* What dose this golang program print?
* Panic
* Does not compile
* 0
* 1
 */
func main() {
	fmt.Println(bailer())
}

func bailer() (n int) {
	defer func() {
		n++
		recover()
	}()
	panic(nil)
}
