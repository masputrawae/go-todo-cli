package utils

import "unicode"

func IsLetter(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := rune(s[0])
	return unicode.IsLetter(r) && (r <= 'Z' || r <= 'z')
}
