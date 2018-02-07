package main

/**
https://twitter.com/inancgumus/status/932526809773723648
 What this func returns? #golang #quiz
*/
func main() {
	print(func(i int) (n int) {
		defer func(f func() int) {
			n = f()
		}(func() int {
			return i
		})
		return
	}(4))
}
