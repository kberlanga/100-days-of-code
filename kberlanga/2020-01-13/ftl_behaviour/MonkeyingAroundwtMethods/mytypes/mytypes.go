package mytypes

import "strings"

// MyInt is a custom version of the `int` type.
type MyInt int
type MyString string
type MySliceString []string
type MySliceInt []int

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}

// MyStringLen returns the lengt of the string
func (s MyString) MyStringLen() int {
	return len(strings.TrimSpace(string(s)))
}

// MultiString returns joined string by " plus "
func (ss MySliceString) MultiString() string {
	return strings.Join(ss, " plus ")
}

// SumInts returns a sum of the []int
func (si MySliceInt) SumInts() int {
	s := 0
	for _, i := range si {
		s += i
	}
	return s
}
