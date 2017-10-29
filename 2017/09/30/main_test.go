package main

import (
	"fmt"
	"testing"
)

/**
 * https://twitter.com/traviscline/status/914005087609020416
 * #golang quiz:
 * https://play.golang.org/p/Pomuyokc-b
 * What output do you get from running this test?
 **/
func TestEcho(t *testing.T) {
	type testCase struct {
		name string
		n    int
	}
	tests := []testCase{
		{"case_1", 1},
		{"case_2", 2},
		{"case_3", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(tt.name, tt.n)
		})
	}
}
