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
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
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

