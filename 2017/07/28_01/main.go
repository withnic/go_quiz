package main

func f(u *uint) { *u++ }

/*
https://twitter.com/MarceloBytes/status/890954798081728512
#golang pop quiz:what does this print?
0
1
Doesn't compile
*/
func main() {
	i := 0
	f((*uint)(&i))
	println(i)
}
