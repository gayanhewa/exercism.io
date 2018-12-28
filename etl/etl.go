package etl

import (
	"strings"
)

// Transform will convert the legacy data to the new format.
func Transform(inputs map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range inputs {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return output
}
