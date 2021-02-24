package main

import (
	"unicode/utf8"
)

// byte count & rune count
func CharCount(s string) (bytes, runes int) {
	bytes = len(s)
	runes = utf8.RuneCountInString(s)
	return
}
