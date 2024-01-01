package string

import (
	"fmt"
	"testing"
)

func TestCountStr(t *testing.T)  {
	got := CountStr("encyclopedia", "e")
	expected := 2

	if got != expected {
		t.Errorf("Expected %q got %q", expected, got)
	}
}

func ExampleCountStr() {
	result := CountStr("cheese", "e")
	fmt.Println(result)
	// Output: 3
}

func BenchmarkCountStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountStr("hello", "l")
	}
}