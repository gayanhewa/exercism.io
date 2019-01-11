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
	letters := map[rune]bool{}
	for _, letter := range strings.ToLower(word) {
		if unicode.IsSpace(letter) || letter == '-' {
			continue
		}
		if letters[letter] {
			return false
		}
		letters[letter] = true
	}
	return true
}
