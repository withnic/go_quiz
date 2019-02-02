package main

import "fmt"

// https://twitter.com/bot_golang/status/1044977904772308992
// #golang quiz
//
//What does the following program print?
// Runtime error
// Compilation error
// 0
func main() {
m := make(map[string]int)
fmt.Println(m["hello"])
}