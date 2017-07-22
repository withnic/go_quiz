package main

import "fmt"

/*
* https://twitter.com/SchumacherFM/status/883272643012685825
* Prints what ?
* 0 0
* 0 1
* 0 32
 */
func main() {

	b1 := []byte{}
	b2 := []byte("")
	fmt.Println(cap(b1), cap(b2))
}
