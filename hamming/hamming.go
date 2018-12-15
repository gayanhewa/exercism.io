package hamming

import (
	"errors"
	"strings"
)

// Distance should return the hamming distance between two patterns, or -1 if not applicable.
func Distance(a, b string) (int, error) {
	a = strings.TrimSpace(a)
	b = strings.TrimSpace(b)

	if a == "" || b == "" || a == b {
		return 0, nil
	}
	if len(a) != len(b) {
		return 0, errors.New("unable to calculate hamming distance")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
