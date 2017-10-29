package main

import "net/url"

/**
 * https://twitter.com/joncalhoun/status/916388460138811392
 * #golang pop quiz - what panic will this code cause?
 */
func main() {
	var v url.Values
	v.Set("name", "")
}
