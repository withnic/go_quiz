package main

import "encoding/json"

/**
 * https://twitter.com/mvdan_/status/978560459388473344
 * sneaky #golang pop quiz:
 */
func main() {
	var b bool
	err := json.Unmarshal([]byte("null"), &b)
	println(b, err != nil)
}
