package anagram

import (
	"strings"
)

func isAnargram(firstWord string, secWord string) bool {
	for _, letterInFirst := range firstWord {
		var exists bool
		for i, letterInSec := range secWord {
			if letterInFirst == letterInSec {
				secWord = secWord[:i] + secWord[i+1:]
				exists = true
				break
			}
		}
		if exists == false {
			return false
		}
	}
	return true
}

func isInvalidStrings(word string, candidate string) bool {
	return word == candidate || len(word) != len(candidate)
}

// Detect determines if the given list contains an anagram of the word.
func Detect(word string, list []string) (anagrams []string) {
	word = strings.ToLower(word)
	for _, candidate := range list {
		if isInvalidStrings(word, strings.ToLower(candidate)) {
			continue
		}
		if isAnargram(strings.ToLower(candidate), word) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
