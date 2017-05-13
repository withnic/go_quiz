package main

import "fmt"

// @avecheney 2017/05/13
func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	<-c
	close(c)
	//  what does this print?
	fmt.Println(<-c)
}
