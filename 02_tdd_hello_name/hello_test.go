package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello("Chris")
	expected := "Hello, Chris"

	if actual != expected {
		t.Errorf("Expected %q, Actual: %q", expected, actual)
	}
}
