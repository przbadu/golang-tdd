package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to peope", func(t *testing.T) {
		actual := Hello("Chris", "")
		expected := "Hello, Chris"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := Hello("Elodie", "es")
		expected := "Hola, Elodie"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in French", func(t *testing.T) {
		actual := Hello("Harry", "fr")
		expected := "Bonjour, Harry"

		assertCorrectMessage(t, actual, expected)
	})

	t.Run("in Nepali", func(t *testing.T) {
		actual := Hello("पुष्प राज", "np")
		expected := "नमस्ते, पुष्प राज"

		assertCorrectMessage(t, actual, expected)
	})
}

func assertCorrectMessage(t testing.TB, actual, expected string) {
	// making this a helper function, so that if test fails
	// it displays line number from the actual function rather
	// than from this helper function.
	t.Helper()

	if actual != expected {
		t.Errorf("Expected %q, Actual: %q", expected, actual)
	}
}
