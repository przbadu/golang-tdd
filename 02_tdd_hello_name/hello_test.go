package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to peope", func(t *testing.T) {
		actual := Hello("Chris")
		expected := "Hello, Chris"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %q, Actual: %q", expected, actual)
	}
}
