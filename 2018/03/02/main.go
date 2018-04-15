package main

import "fmt"

/**
 * https://twitter.com/variadico/status/969469151700910081
 * 	What does this print?
 */
func main() {
	m := map[string]interface{}{"a": "b"}
	fmt.Printf("type=%[1]T; val=%#[1]v\n", m["a"])
}
