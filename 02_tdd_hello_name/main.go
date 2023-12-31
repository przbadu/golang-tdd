package main

import "fmt"

const (
	spanish            = "es"
	french             = "fr"
	nepali             = "np"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	nepaliHelloPrefix  = "नमस्ते, "
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
	case nepali:
		prefix = nepaliHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
