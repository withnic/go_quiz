package main

import (
	"fmt"
	"time"
)

/**
https://twitter.com/davecheney/status/444707522591477760
#Q. golang pop quiz: time.Duration is defined as int64. What would happen if it were declared uint64 ? https://golang.org/pkg/time/#Duration
**/

/**
rf‏ @rf  2014年3月15日
@davecheney Single value couldn't represent either "in 5 seconds" or "5 seconds ago," but I get the feeling you mean something subtler.
**/

/**
Hiroaki Nakamura‏ @hnakamur2  2014年3月15日
@davecheney You need func (t Time) Sub(d Duration) Time as well as Add and rename (t Time) Sub(u Time) Duration to something else.
**/
func main() {
	x := -1 * time.Second
	fmt.Print(x)
}
