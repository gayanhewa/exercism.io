// Package bob implements functions that help us infer a teenagers expressions.
package bob

import (
	"strings"
	"unicode"
)

func hasLetters(remark string) bool {
	hasLetters := false
	for _, value := range []rune(remark) {
		if unicode.IsLetter(value) {
			hasLetters = true
			break
		}
	}
	return hasLetters
}

func isYelling(remark string) bool {
	return hasLetters(remark) && strings.ToUpper(remark) == remark
}

func isQuestion(remark string) bool {
	return strings.LastIndex(remark, "?") == (len(remark) - 1)
}

// Hey infers the response based on the remark given to a teenager.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	hasLetters := hasLetters(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isQuestion(remark) && (!isYelling(remark) || !hasLetters):
		return "Sure."
	case isYelling(remark) && !isQuestion(remark):
		return "Whoa, chill out!"
	case isYelling(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	default:
		return "Whatever."
	}
}
