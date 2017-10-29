package main

/**
* https://twitter.com/SamWhited/status/910887142708170754
* #golang pop quiz, what does this do:
*
* Does not build
* Prints "nill"
 */

func main() {
	var err error
	switch err.(type) {
	case nil:
		print("nil")
	}
}
