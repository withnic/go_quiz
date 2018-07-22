package main

import "fmt"

func four(s []int) {
	s = append(s, 4)
}

/**
* https://twitter.com/davecheney/status/1019359774264082432
* what does this program print?
* [1 2 3 4]
* [4]
* [1 2 3]
* doesn't compile
 */
func main() {
	s := []int{1, 2, 3}
	four(s)
	fmt.Println(s)
}
