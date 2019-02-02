package main

import (
"fmt"
)
// https://twitter.com/bot_golang/status/1037351913853841408
// What does the following program print?
// 5
// Hello5
// Compile error
func main() {
	v, _ := fmt.Print("Hello")
	fmt.Print(v)
}
