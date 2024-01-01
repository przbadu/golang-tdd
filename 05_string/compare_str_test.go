package string

import "testing"

func TestCompareStr(t *testing.T) {
	t.Run("compare if two strings are equal", func(t *testing.T) {
		got := CompareStr("hello", "hello")
		expected := true

		if expected != got {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	})
}