package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
#golang lunch time quiz: What will this print?
[1]
[1,2,3,4,5,6,7,8,9]
err
panic
https://twitter.com/StabbyCutyou/status/877936095039803392
*
*/
func main() {
	js := `[1,2,3,4,5,6,7,8,9]`
	s := make([]int, 0, 1)
	err := json.Unmarshal([]byte(js), &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
