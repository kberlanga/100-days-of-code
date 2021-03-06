package mytypes_test

import (
	"mytypes"
	"testing"
)

func TestMyString_MyStringLen(t *testing.T) {
	tests := []struct {
		name string
		s    mytypes.MyString
		want int
	}{
		{
			name: "normal string",
			s:    "Hello, Gophers!",
			want: 15,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "other string",
			s:    "3.. 2.. 1.. Go!",
			want: 15,
		},
		{
			name: "other empty string",
			s:    "         ",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MyStringLen(); got != tt.want {
				t.Errorf("MyString.MyStringLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySliceString_MultiString(t *testing.T) {
	tests := []struct {
		name string
		ss   mytypes.MySliceString
		want string
	}{
		{
			name: "multi string",
			ss:   []string{"a", "b", "c"},
			want: "a plus b plus c",
		},
		{
			name: "empty multi string",
			ss:   []string{},
			want: "",
		},
		{
			name: "multi string",
			ss:   []string{"a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ss.MultiString(); got != tt.want {
				t.Errorf("MySliceString.MultiString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySliceInt_SumInts(t *testing.T) {
	tests := []struct {
		name string
		si   mytypes.MySliceInt
		want int
	}{
		{
			name: "normal sum",
			si:   []int{1, 2, 3},
			want: 6,
		},
		{
			name: "zero",
			si:   []int{-1, -2, 3},
			want: 0,
		},
		{
			name: "normal sum",
			si:   []int{10, 876, -56},
			want: 830,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.si.SumInts(); got != tt.want {
				t.Errorf("MySliceInt.SumInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
