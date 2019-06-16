package main

// https://twitter.com/tenntenn/status/1130028032423391232
// #golang quiz 3
// compile error
// 9
// 10
// 11
func main() {
	ns := []int{010: 200, 005: 100}
	print(len(ns))
}
