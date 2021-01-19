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

func TestStringUppercaser_ToUpper(t *testing.T) {
	su := mytypes.StringUppercaser{}
	su.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()

	if want != got {
		t.Errorf("StringUppercaser.ToUpper() = %v, want %v", got, want)
	}
}

func TestSquare(t *testing.T) {
	p := 4
	want := 16
	mytypes.Square(&p)
	got := p

	if want != got {
		t.Errorf("Square() = %v, want %v", got, want)
	}
}

func TestSwapInts(t *testing.T) {
	x := 4
	y := 8
	want_x := 8
	want_y := 4
	mytypes.SwapInts(&x, &y)

	if want_x != x || want_y != y {
		t.Errorf("SwapInts() = (%v, %v), (want_x %v, want_y %v)", x, y, want_x, want_y)
	}
}

func TestDouble(t *testing.T) {
	got := mytypes.MyInt(4)
	got.Double()
	want := mytypes.MyInt(8)

	if got != want {
		t.Errorf("Double() = %v, want %v", got, want)
	}
}

func TestMultiplyBy(t *testing.T) {
	p := mytypes.MyInt(3)
	got := mytypes.MyInt(4)
	got.MultiplyBy(&p)
	want := mytypes.MyInt(12)

	if got != want {
		t.Errorf("MultiplyBy(%v) = %v, want %v", p, got, want)
	}
}
