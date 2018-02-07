package main

/**
#golang pop quiz, what does calling the following function return:
0
123
panic
Does not compile
*/
func main() {
	foo()
}

func foo() (i int) {
	defer func() {
		println(i)
	}()
	return 123
}
