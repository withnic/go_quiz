package main

/**
https://twitter.com/bradfitz/status/951534740405829632
#golang pop quiz (@davecheney style): what does this program print?
* 4
* 3
* depends
* doesn't compile
*/
func main() {
	print(len(map[interface{}]int{
		new(int):      1,
		new(int):      2,
		new(struct{}): 3,
		new(struct{}): 4,
	}))
}
