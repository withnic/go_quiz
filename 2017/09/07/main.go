package main

import "fmt"

/**
https://twitter.com/importantshock/status/905565924446670848
fun Golang quiz: what does (+(-1.0)) evaluate to?
*/
func main() {

	a := (+(-1.0))

	fmt.Print(a)
}
