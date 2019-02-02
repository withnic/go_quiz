package main

func bar(i int) {
	println(i)
}

func foo() {
	i := 0
	defer bar(i)
	i = 1
}

// https://twitter.com/hajimehoshi/status/1084710551756800000
// Go Quiz: What is the output?
//0
//1
//🤔
func main() {
	foo()
}
