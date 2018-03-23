package main

import "fmt"

/**
https://twitter.com/fzerorubigd/status/966646995669667840
#Golang pop quiz
* 2 , 2222
* 2222 , 2
* 2 , 2
* 2222 , 2222
*/
func main() {
	var a = make([]int, 0, 10)
	a = append(a, 1, 2, 3, 4)
	b := append(a, 5, 6, 7)
	c := append(a, 8, 9, 10, 111, 12, 13, 14)
	a[1] = 2222
	fmt.Println(b[1], ",", c[1])
}
