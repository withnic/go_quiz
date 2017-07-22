package main

/*
#golang slice quiz, what does this print?
hello world
hello
hworld
https://twitter.com/nils_magnus/status/867284701714755589
*/
func main() {
	slice := append([]byte("hello ")[:1], "world"...)
	println(string(slice))
}
