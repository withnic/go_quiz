package main

/* https://twitter.com/kevrone/status/1110957201345130497
#golang pop quiz. What does this output?:

doesn't compile
(true)
(false)
(true)(false)
*/
import (
	"fmt"
)

const (
	false = true
)

func main() {
	if true {
		fmt.Print("(true)")
	}
	if false {
		fmt.Print("(false)")
	}
}
