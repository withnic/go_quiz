package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/**
https://twitter.com/shurcooL/status/903465910110785541
#golang pop quiz
Logically, are these equivalent?
	err := json.Unmarshal(data, &v)
	err = json.NewDecoder(bytes.NewReader(data)).Decode(&v)

Same Logical behavior
Different behavior (how?)
*/

var data = ([]byte)("[1,2,3]\n{\"extra\": \"obj2\"")

/**
https://twitter.com/9sqweek/status/903522373831344132
*/
func main() {
	var v1, v2 interface{}
	err1 := json.Unmarshal(data, &v1)
	err2 := json.NewDecoder(bytes.NewReader(data)).Decode(&v2)
	fmt.Println(err1)
	fmt.Println(v1)
	fmt.Println(err2)
	fmt.Println(v2)
}
