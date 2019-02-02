package main

import "fmt"

// https://twitter.com/mlowicki/status/1039993406364237830
// #golang quiz:
// what this program does?
// 1
// 1, 2, 3
// 1, 2, 3, panic
// panic(deadlock)
func main() {
	ch := make(chan int, 3)
	ch<-1
	ch<-2
	ch<-3
	for {
		select {
		case n := <-ch:
			fmt.Println(n)
			break
		}
	}
}