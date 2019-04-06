package main

/* https://twitter.com/bot_golang/status/1101181308259237889
#golang quiz

What does the following program print?

Complication Error
A
Runtime Error
*/
import (
	"fmt"
)

func main() {
	i := 65
	s := string(i)
	fmt.Println(s)
}
