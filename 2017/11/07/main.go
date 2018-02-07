package main

/**
https://twitter.com/davecheney/status/927733411070074880
#golang pop quiz: does this program compile?
*/
func ch(ch chan int) {
        ch <--<-ch
}

func main() {} 