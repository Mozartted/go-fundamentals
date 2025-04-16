package depsinjection

import (
	"fmt"
	"io"
	"net/http"
)

// Interfaces are the key to golang's Dependency injections
// both io.Writer and http.ResponseWriter both implement a Write
// function which takes in []byte as arguements
func Greet[T string](b io.Writer, name T) {
	fmt.Fprintf(b, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
