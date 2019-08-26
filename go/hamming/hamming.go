package hamming

import (
	"errors"
)

// Distance should return the hamming distance between two patterns, or -1 if not applicable.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("unable to calculate hamming distance")
	}

	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
