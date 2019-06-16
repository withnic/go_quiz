package main

// #golang quiz 1
// https://twitter.com/tenntenn/status/1130026890092109826
// compile error
// panic
// AA
// AAAA
func init() {
	main()
}

func init() {
	panic()
}

var panic = main

func main() {
	print("A")
}

func init() {
	main()
}
