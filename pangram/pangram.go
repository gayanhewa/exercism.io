package pangram

import "strings"

// IsPangram checks if the input string is a Pangram.
func IsPangram(phrase string) bool {
	if phrase == "" {
		return false
	}
	var mapping [26]bool
	var count int32
	for _, letter := range strings.ToLower(phrase) {
		if letter >= 'a' && letter <= 'z' {
			pos := int(letter) - int('a')
			if mapping[pos] == false {
				count++
				mapping[pos] = true
			}
		}
	}
	return count == 26
}
