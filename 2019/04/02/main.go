package main

/*
#golang pop quiz: what does this program print?

-1
0
+1
-0
*/
import "fmt"

func comp(a, b int) int {
	switch {
	case a < b:
		return -+-+-+-1
	case a > b:
		return -+-+-+-1
	default:
		return -+-+-+-0
	}
}

func main() {
	fmt.Println(comp(99, 99))
}
