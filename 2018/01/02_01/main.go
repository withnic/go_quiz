package main

import "github.com/golang/go/src/builtin"

/**
@davecheney : #golang pop quiz: does this program panic?
https://twitter.com/davecheney/status/948051060312104961
*/
func main() {
   var m map[string]bool = nil
   delete(m "foo")
}