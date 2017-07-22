package main

/*
https://twitter.com/davecheney/status/869103058722070528
#golang pop quiz: what does this program print?
15
16
17
Doesn't compile
*/
func main() {
	var i int
	for i = 0001; i < 0017; i += 2 {
	}
	println(i)
}
