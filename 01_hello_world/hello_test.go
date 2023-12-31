package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello()
	expected := "Hello, World"

	if actual != expected {
		t.Errorf("Actual %q expected %q", actual, expected)
	}
}
