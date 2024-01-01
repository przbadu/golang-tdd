package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T)  {
	numbers := [5]int{1, 2 ,3 ,4 ,5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("Expected %d, got %d, given %v", expected, got, numbers)
	}
}

func ExampleSum() {
	numbers := [5]int{1, 4, 5, 10, 20}
	result := Sum(numbers)
	fmt.Println(result)
	// Result: 40
}