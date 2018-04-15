package main

import "fmt"

type myByte byte

func (b myByte) String() string {
	return fmt.Sprintf("%x", b)
}

/**
 * https://twitter.com/LobaroHH/status/978375616096538624
 *  #GoLang pop quiz
 * What is the output of the following program?
 * 1
 * 01
 * 31
 * stack overflow
 */
func main() {
	b := myByte(1)
	fmt.Printf("%x", b)
}
