package main

import "fmt"

const (
	spanish            = "es"
	french             = "fr"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix // initialize variable

	switch lang {
	case spanish:
		prefix = spanishHelloPrefix // assign different value
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
