package main

import (
	"fmt"
	"net/http"
)

/**
https://twitter.com/kevrone/status/961741465025241089
#golang pop quiz. What content type does this detect?
* text/plain
* application/json
* text/html
* application/octet-stream
*/
func main() {
	js := []byte(`[{}]`)
	contentType := http.DetectContentType(js)
	fmt.Println(contentType)
}
