package main

/*
#golang pop quiz: what does this program print?
3 5
8 3
8 8
Infinite loop
https://twitter.com/davecheney/status/869376332429352960
*/
func main() {
	a, b := 3, 5
	a, b = a+b, a
	println(a, b)
}
