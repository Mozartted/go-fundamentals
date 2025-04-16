package main

import (
	// depsinjection "fundamental/deps_injection"
	"fundamental/depsinjection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depsinjection.MyGreeterHandler)))
}
