// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey method return a specifc phrase depending on the input
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if remark == "" {
		return "Fine. Be that way!"
	}
	isNumeric := isNumeric(remark)
	isUpper := strings.ToUpper(remark) == remark
	isQuestion := string(remark[len(remark)-1]) == "?"
	if isUpper && isQuestion && !isNumeric {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isUpper && !isNumeric {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}

func isNumeric(str string) (numeric bool) {
	numeric = true
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			numeric = false
		}
	}
	return numeric
}
