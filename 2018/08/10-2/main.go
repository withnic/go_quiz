package main

import "fmt"

// https://twitter.com/bot_golang/status/1027635332039106560
// #golang quiz
//What does the following program print?
// Complication Error
// hello world
type person struct {
firstName string
}

func (p *person) test() {
fmt.Println("hello world")
}

func main() {
var p *person = nil
p.test()
}