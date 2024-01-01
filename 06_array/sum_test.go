package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T)  {
	t.Run("sum of collection of n numbers", func (t *testing.T)  {
		numbers := []int{1, 2 ,3 ,4 ,5}

		got := Sum(numbers)
		expected := 15
	
		if got != expected {
			t.Errorf("Expected %d, got %d, given %v", expected, got, numbers)
		}	
	})
}

func TestSummAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSum() {
	numbers := []int{5, 10}
	result := Sum(numbers)
	fmt.Println(result)
	// Result: 15
}

func ExampleSumAll()  {
	result := SumAll([]int{2, 3}, []int{5, 5})
	fmt.Println(result)
	// Output []int{5, 10}
}