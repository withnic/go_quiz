package main

/*
https://twitter.com/davecheney/status/896701753365626881
#golang pop quiz
prints 8
prints 9
prints something else
doesn't compile
*/
func main() {
	type P *int
	type Q *int
	var p P = new(int)
	*p += 8
	var x *int = p
	var q Q = x
	*q++
	print(*p)
}
