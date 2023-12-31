package main

import (
	"fmt"
)

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

	return greetingPrefix(lang) + name
}

// (prefix string): is a named return value
// this will create a variable called `prefix`
// default value is 0 for int, "" for string.
// All you need to do is call `return` at the end of your function.
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case nepali:
		prefix = nepaliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
