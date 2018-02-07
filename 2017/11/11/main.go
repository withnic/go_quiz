package main

import "fmt"

/**
https://twitter.com/StabbyCutyou/status/929097635293749248
#golang pop quiz: what does this print?
0
1
2
panic
*/
func main() {
	i := 0
	f := func() int {
		i++
		return i
	}
	c := make(chan int, 1)
	c <- f()
	select {
	case c <- f():
	default:
		fmt.Println(i)
	}
}
