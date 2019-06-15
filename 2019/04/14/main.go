package main

import "fmt"

// #golang pop quiz: what does this program print?
// https://twitter.com/davecheney/status/1117216793309765632
// 0
// 1
// -1
// dose not compile
func f(a, b int) int {
	if a > b {
		a = b
	} else if a < b {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(f(64, 1))
}
