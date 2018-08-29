package main

import "fmt"

//#golang quiz
//https://twitter.com/bot_golang/status/1026859630427660288
//What does the following program print?
// George Eliot
// Complication Error
// Eliot
type person struct {
firstName string
}

func main() {
p := person{}
p.firstName, lastName := "George", "Eliot"
fmt.Println(p.firstName, lastName)
}