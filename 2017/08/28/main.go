//https://twitter.com/SamWhited/status/902019104307077120
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < '丠'; i++ {
		f, _ := os.Create(fmt.Sprint(i) + ".go")
		f.WriteString("package a")
		f.Close()
	}
}
