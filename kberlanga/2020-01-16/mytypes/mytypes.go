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

// MyBuilder is a custom version of
// strings.Builder.
// type MyBuilder strings.Builder

// MyBuilder wraps a strings.Builder.
type MyBuilder struct {
	Contents strings.Builder
}

// Hello returns the string "Hello, Gophers!".
func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

// StringUppercaser wraps a strings.Builder.
type StringUppercaser struct {
	strings.Builder
}

// Hello returns the string to uppercase.
func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.String())
}

// takes a pointer to int, returns nothing,
// and has the side‐effect of multiplying the original int value by itself.
func Square(p *int) {
	*p *= *p
}

func SwapInts(x, y *int) {
	*x, *y = *y, *x
}

func (i *MyInt) Double() {
	*i *= 2
}

func (i *MyInt) MultiplyBy(f *MyInt) {
	*i *= *f
}
