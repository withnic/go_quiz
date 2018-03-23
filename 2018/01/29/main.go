package main

import (
	"fmt"
)

/**
https://twitter.com/davecheney/status/957712514942251008
#golang pop quiz, what does this program print?
* 3 then 5
* 3 then 3
* nothing (syntax error)
* 3 3 3 ... (infinite loop)
*/
func main() {
	var i int
	for i < 5 {
		i = 5
		i := 3
		fmt.Println(i)
	}
	fmt.Println(i)
}
