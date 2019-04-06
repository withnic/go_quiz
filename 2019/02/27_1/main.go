package main

/* https://twitter.com/tvii/status/1100761746074161152
#golang popup quiz time:
What will the following code print?

main.FooType ""
main.FooType foot
syntax error
other versions
*/
import "fmt"

type FooType string
const (
	Foo FooType = "foo"
	Bar
)
func main() {
	fmt.Printf("%T %s", Bar, Bar)
}
