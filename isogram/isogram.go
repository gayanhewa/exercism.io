package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram check's if the given word is an isogram, with spaces and hyphens exepted
func IsIsogram(word string) bool {
	if word == "" {
		return false
	}
	letters := map[rune]rune{}
	for _, letter := range strings.ToLower(word) {
		if unicode.IsSpace(letter) || letter == '-' {
			continue
		}
		if letters[letter] != 0 {
			return false
		}
		letters[letter] = letter
	}
	return true
}
