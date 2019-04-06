package main

/*
https://twitter.com/DudeWhoCode/status/1104701111087247360
#Golang quiz time:

Whats the output and why?

5 4 3 2 1
1 2 3 4 5
0 0 0 0 0
5 5 5 5 5
*/
import "fmt"

func main() {
	foo()
}

func foo() {
	for i := 5; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
}
