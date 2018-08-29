package main

import (
"fmt"
"runtime"
"sync"
)

// https://twitter.com/davecheney/status/1031699003803463680
// #golang pop quiz: what does this program print?
// Hello, playground
// It panics
// nothing, it does'nt compile
func main() {
var wg sync.Waitgroup
wg.Add(1)
go func() {
defer wg.Done()
runtime.Goexit()
}()
wg.Wait()
fmt.Println("Hello, playground")
} 