package string

import "strings"

func CompareStr(left, right string) bool {
	return strings.Compare(left, right) == 0
}

func ContainsStr(left, right string) bool {
	return strings.Contains(left, right)
}

func CountStr(word, char string) int {
	return strings.Count(word, char)
}