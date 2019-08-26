package wordcount

import (
	"log"
	"regexp"
	"strings"
)

// Frequency type.
type Frequency map[string]int

func clean(s string) []string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	hasCommas := strings.Contains(s, ",")
	hasSpaces := strings.Contains(s, " ")
	if hasCommas && !hasSpaces {
		return strings.Split(s, ",")
	}
	if hasSpaces && !hasCommas {
		return strings.Split(s, " ")
	}
}

// WordCount returns a map of Word frequencies in the given phrase.
func WordCount(phrase string) Frequency {
	f := make(Frequency)
	for _, word := range strings.Split(phrase, " ") {
		f[word]++
	}

	return f
}
