package main

import (
	"fmt"
	"time"
)

/*
#golang pop quiz: what will this print?
https://twitter.com/peterbourgon/status/865159346149314560
*/
func main() {
	var i int
	go fmt.Println(i)
	i++
	go func() { fmt.Println(i) }()
	i++
	go func(x int) { fmt.Println(x) }(i)
	i++
	time.Sleep(time.Second)
}
