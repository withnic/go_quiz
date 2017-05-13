package main

import "testing"

var result []byte

/**
@davecheney
https://twitter.com/davecheney/status/364240081223024640
pop quiz number #2. Which of these three is faster ? Try to guess before running the code.
**/

func BenchmarkMakeOneByte(b *testing.B) {
	var v []byte
	for i := 0; i < b.N; i++ {
		v = make([]byte, 1)
	}
	result = v
}

func BenchmarkOneLiteralByte(b *testing.B) {
	var v []byte
	for i := 0; i < b.N; i++ {
		v = []byte{0}
	}
	result = v
}

var array [1]byte

func BenchmarkSliceOneByte(b *testing.B) {
	var v []byte
	for i := 0; i < b.N; i++ {
		v = array[:1]
	}
	result = v
}
