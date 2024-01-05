package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	assert(t, expected, got)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Area(rectangle)
	expected := 72.0

	assert(t, expected, got)
}

func assert(t testing.TB, expected, got float64) {
	t.Helper()

	if got != expected {
		t.Errorf("Expected %.2f but got %.2f", expected, got)
	}
}
