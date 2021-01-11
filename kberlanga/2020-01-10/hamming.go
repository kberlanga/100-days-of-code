package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("long different strands")
	}

	a1 := []rune(a)
	b1 := []rune(b)
	e := 0
	for i := 0; i < len(a); i++ {
		if a1[i] != b1[i] {
			e++
		}
	}

	return e, nil
}
