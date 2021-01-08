// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate function receive a string to convert that to acronym.
func Abbreviate(s string) string {
	words := strings.Split(s, " ")
	acronym := ""
	for _, word := range words {
		// flag := true
		for i, letter := range word {
			ascii := int(letter)
			if ascii > 64 && ascii < 91 || ascii > 96 && ascii < 123 {
				if i == 0 {
					acronym += strings.ToUpper(word[:1])
				}
			} else {
				if len(word) > 1 && i < len(word)-1 && ascii != 39 {
					test := []rune(word)
					acronym += strings.ToUpper(string(test[i+1]))
				}
			}
		}
	}
	return acronym
}
