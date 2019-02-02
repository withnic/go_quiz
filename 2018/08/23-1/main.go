package main

import "fmt"

// https://twitter.com/bot_golang/status/1032298657884696583
//#golang quiz #goto #statement
//What does the following program print?
// Complication Error
// n <= 4;
// n > 4;
// n <= 4;n > 4;
func main() {
n := 3
if n > 4 {
goto Last
}
fmt.Print("n <= 4;")
Last:
fmt.Print("n > 4;")
}
