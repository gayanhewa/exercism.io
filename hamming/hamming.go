package hamming

import (
	"bytes"
	"errors"
)

// Distance should return the hamming distance between two patterns, or -1 if not applicable.
func Distance(a, b string) (int, error) {
	aString := bytes.TrimSpace([]byte(a))
	bString := bytes.TrimSpace([]byte(b))

	if len(aString) != len(bString) {
		return -1, errors.New("Unable to calculate hamming distance.")
	}

	if bytes.Equal(aString, bString) {
		return 0, nil
	}

	distance := 0

	for index, value := range aString {
		if value != bString[index] {
			distance++
		}
	}

	return distance, nil
}
