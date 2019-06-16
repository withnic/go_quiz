package main

import (
	"fmt"
)

type hello struct{}

func (h hello) data() {
	fmt.Println("hello")
}

type hi hello

/*
https://twitter.com/bot_golang/status/1126467123117563905
#golang quiz
What does the following program print?

*/
func main() {
	hi.data()
}
