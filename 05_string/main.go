package string

import "strings"

func CompareStr(left string, right string) bool {
	return strings.Compare(left, right) == 0
}