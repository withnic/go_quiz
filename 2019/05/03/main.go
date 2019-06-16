package main

import (
	"fmt"
)

/*
https://twitter.com/bot_golang/status/1124018542444208130
#golang quiz
What does the following program print?

world
panics at runtime
*/
func main() {
	var a map[string]string
	a["hello"] = "world"
	fmt.Println(a["hello"])
}
