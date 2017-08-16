package main

/*
https://twitter.com/davecheney/status/893801568796950528
#golang pop quiz: does this program

loop forever
exit cleanly
not compile
*/
func main() {
again:
	for i := 0; i < 10; i++ {
		if i > 6 {
			break again
		}
	}
}
