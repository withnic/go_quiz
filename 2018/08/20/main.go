package main

import "fmt"

// https://twitter.com/davecheney/status/1031389514890006528
// #golang pop quiz: what does this program print?
// Bueno
// Hola
// Nada, it won't compile
type Q struct{}

func (Q) Hola() string { return "Bueno" }

func main() {
var q Q
fmt.Println(Q.Hola(q))
}
