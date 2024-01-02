package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("Expected %d, got %d, given %v", expected, got, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper() // helper function
		if !slices.Equal(got, want) {
			t.Errorf("got %v expected %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func ExampleSum() {
	result := Sum([]int{5, 10})
	fmt.Println(result)
	// Result: 15
}

func ExampleSumAll() {
	result := SumAll([]int{2, 3}, []int{5, 5})
	fmt.Println(result)
	// Output {5, 10}
}

func ExampleSumAllTails() {
	result := SumAllTails([]int{1, 2}, []int{1, 2, 3, 10})
	fmt.Println(result)
	// Output {2, 15}
}

