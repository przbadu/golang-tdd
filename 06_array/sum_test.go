package main

import (
	"fmt"
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

func ExampleSum() {
	numbers := []int{5, 10}
	result := Sum(numbers)
	fmt.Println(result)
	// Result: 15
}