package main

import "fmt"

/**
 * #golang pop quiz: What is the final state of a?
 *  https://twitter.com/joncalhoun/status/974701114208739328
 */
func main() {
	a := []int{1, 2, 3, 4}
	b := variadic(a...)
	b[0], b[1] = b[1], b[0]
	fmt.Println("a=", a)
}

func variadic(ints ...int) []int {
	return ints
}
