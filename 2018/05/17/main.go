package main

import (
	"fmt"
	"time"
)

/**
* https://twitter.com/inancgumus/status/996789929928724486
* What is the output of this golang program?
// hint 1s = 1.000.000.000ns
* Panic
* Error
* 1s
* 2s
*/
func main() {
	fmt.Println(time.Duration(1e9) + 1e9)
}
