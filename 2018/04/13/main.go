package main

import "fmt"

/**
* https://twitter.com/SamWhited/status/984487196752252928
* what dose foo() return?
*
* 0, "baz"
* 4, "baz"
* Does not compile
 */
func main() {
	fmt.Println(foo())
}

func baz(i *int) string {
	*i = 4
	return "baz"
}

func foo() (int, string) {
	bar := 0
	return bar, baz(&bar)
}
