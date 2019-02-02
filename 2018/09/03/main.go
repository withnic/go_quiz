package main

import (
"fmt"
"time"
)

// https://twitter.com/bot_golang/status/1036598781733523458
// #golang quiz
//
//What does the following print?
func hello() {
time.Sleep(10 * time.Second)
fmt.Println("go")
}

func main() {
go hello()
}
