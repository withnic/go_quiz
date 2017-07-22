package main

import (
	"fmt"
)

/*
Find a bug? Make a #popquiz!
What is the value of Total right before the program exits?
*https://twitter.com/JF0LKINS/status/875558826321358848
*/
func main() {

	var Total int

	for i := 0; i <= 10; i++ {
		Total = Total + i
	}
	fmt.Println(Total)

	Total = 00 + 01 + 02 + 03 + 04 + 05 + 06 + 07 + 8 + 9 + 010

}
