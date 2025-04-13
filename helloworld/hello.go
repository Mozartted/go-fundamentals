package main

import "fmt"

const (
	spanish             = "Spanish"
	englishHelloPrefix  = "Hello, "
	spanishhHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		return "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "english":
		prefix = englishHelloPrefix
	case "spanish":
		prefix = spanishhHelloPrefix
	}
	return
}

func main() {
	fmt.Print(Hello("Tony", "spanish"))
}
