package main

import (
	"fmt"
	"strings"
)

// https://twitter.com/mueslix/status/1131003355449442304
// #golang pop quiz: what does this code print out?
// 0
// 1
// Undefined output
func main() {
	s := strings.Split("", "@")
	fmt.Println(len(s))
}
