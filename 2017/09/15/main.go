package main

import (
	"sync"
)

/**
*https://twitter.com/davecheney/status/908454851046076417
Sure,that compiles
Nah, that won't compile
Hey, wait a minute :thinking_face:
**/

type mu = sync.Mutex

type T struct {
	mu
}

func main() {
	t := T{}
	t.Lock()
}
