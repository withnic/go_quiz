package main

func f(u *uint) { *u++ }

/*
https://twitter.com/davecheney/status/890783014178799616
#golang pop quiz:what does this print?
0
1
Doesn't compile
*/
func main() {
	i := 0
	f(&(uint)(i))
	println(i)
}
