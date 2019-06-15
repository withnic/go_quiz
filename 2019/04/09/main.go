package main

import "context"

// golang pop quiz: what does this do?
// https://twitter.com/mvdan_/status/1115300377434443776

func f() (cancel func()) {
	_, cancel = context.WithCancel(context.TODO())
	cancel2 := func() {
		cancel()
		println("cancelled")
	}
	return cancel2
}

func main() {
	f()()
}
