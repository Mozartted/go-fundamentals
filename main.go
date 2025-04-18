package main

import (
	// depsinjection "fundamental/deps_injection"
	// "fundamental/depsinjection"
	// "log"
	// "net/http"
	"fundamental/mock"
	"os"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(depsinjection.MyGreeterHandler)))
	mock.Countdown(os.Stdout, &mock.DefaultSleeper{})
}
