package main

import "time"

/*
#golang pop quiz: does this program
exit successfully
deadlock
dude, it doesn't compile
https://twitter.com/davecheney/status/875216084906909696
*/
func main() {
	t := time.NewTicker(100)
	for range t.C {
		t.Stop()
	}
}
