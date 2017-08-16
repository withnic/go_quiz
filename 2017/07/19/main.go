package main

import "fmt"

/*
https://twitter.com/janiszt/status/887583428924895232
#golang pop quiz

[0 1 2 3][0 1 2 3]
[0 1 2 3][0 1 2 4]
[0 1 2 4][0 1 2 3 4]
*/
func main() {
	x := []int{0, 1}
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}
