package main

// https://twitter.com/empijei/status/1098570740729868290
// #golang #pop #quiz calling "crasher":
//
// Panics
// Does not panic
// This doesn't compile
type mb string


func main() {
	crasher()
}

func crasher() {
	defer func() { recoverer() }()
	panic(mb("a"))
}

func recoverer() {
	if r := recover(); r != nil {
		if _, ok := r.(mb); ok {
			return
		}
		panic(r)
	}
}
