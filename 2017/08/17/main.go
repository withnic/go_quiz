package main

/*
https://twitter.com/joeshaw/status/898200517343670272
in @davecheney style, #golang quiz:
what does this print, and why?
hi hello greetings yo
hi hello greetings
something else
it panics
*/
func main() {
	l := []string{"hi", "hello", "greetings"}
	for _, x := range l {
		println(x)
		if x[0] == 'g' {
			l = append(l, "salutations")
		}
	}
}
