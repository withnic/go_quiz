package main

// https://twitter.com/tenntenn/status/1130027548404936704
// #golang quiz 2
// compile error
// panic
// 0
// 1
func main() {
	a := &[...]int{1}
	println(a[cap(a)-len(a)])
}
