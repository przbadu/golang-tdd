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

	if lang == spanish {
		return spanishHelloPrefix + name
	}

	if lang == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
