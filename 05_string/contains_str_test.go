package string

import (
	"fmt"
	"testing"
)

func TestContainsStr(t *testing.T) {
	t.Run("'Hello World' should contain 'World'", func (t *testing.T)  {
		got := ContainsStr("Hello World", "World")
		expected := true

		if expected != got {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	})
}

func ExampleContainsStr() {
	result := ContainsStr("seafood", "foo")
	fmt.Println(result)
	// Output: true
}

func BenchmarkContainsStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsStr("seafood", "foo")
	}
}