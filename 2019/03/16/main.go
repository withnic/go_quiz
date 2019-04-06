package main

/*
https://twitter.com/davecheney/status/1106688138007592961
#golang pop quiz: what does this program print?
4
0
nothing program hangs
panic: closed channel
*/
import (
	"fmt"
)

func main() {
	c := make(chan int, 5)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	<-c
	<-c
	<-c
	fmt.Println(<-c)
}
