package main

import "fmt"

/*
#golang pop quiz:
alphabet
battery
cable
none of the above
https://twitter.com/davecheney/status/865062709057826816
*/
func main() {
	var w string
	for _, s := range []string{"alphabet", "battery", "cable"} {
		if s < w {
			w = s
		}
	}
	fmt.Println(w)
}
