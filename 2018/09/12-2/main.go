package main

import (
"fmt"
)

// https://twitter.com/davecheney/status/1039997464361623552
//#golang pop quiz: what does this program print?
//7
//14
//n
//doesn't compile

func main() {
	fmt.Println(string('7'<<1))
}
