package main

import "fmt"

// Separate the concerns of the code
// So that it is easier to test the code
func Hello() string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello())
}
