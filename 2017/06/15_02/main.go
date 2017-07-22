package main

import "fmt"

/*
#golang pop quiz
[1 2 3]
[1 2 3 1 2 3]
[1 2 3 0 1 2]
it't an infinite loop
https://twitter.com/davecheney/status/875200590619570176
*/
func main() {
	x := []int{1, 2, 3}
	for i := range x {
		x = append(x, i)
	}
	fmt.Println(x)
}
