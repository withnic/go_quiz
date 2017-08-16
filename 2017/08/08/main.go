package main

import "fmt"

/*
https://twitter.com/SamWhited/status/894937031754870784
#golang pop quiz what does this print:

map[quiz:value]
map[key:value]
Dose not compile
*/
func main() {
	key := "quiz"
	fmt.Println(map[string]string{
		key: "value",
	})
}
