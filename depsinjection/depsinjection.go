package depsinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet[T string](b io.Writer, name T) {
	fmt.Fprintf(b, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
