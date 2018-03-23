package main

/**
https://twitter.com/blockloop/status/948628738941571074
#golang quiz. Does this print "Hello?" /cc @davecheney
Yes
No
*/
import (
	"fmt"
)

func main() {
	defer fmt.Println("Hello")
	panic("bye bye")
}
