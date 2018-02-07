package main

import (
	"fmt"
	"log"
	"strings"
)

/**
https://twitter.com/davecheney/status/946905445393874944
#golang pop quiz: what does the following program log?
*/
func main() {
	r := strings.NewReader("")
	defer fmt.Println("done")
	_, err := r.Read(make([]byte, 1))
	if err != nil {
		log.Fatal(err)
	}
}
