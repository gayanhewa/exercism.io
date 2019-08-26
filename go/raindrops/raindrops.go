package raindrops

import (
	"strconv"
)

// Convert defines the method that converts numbers into raindrop speak.
func Convert(i int) string {
	var output string
	factors := []struct {
		divisor int
		result  string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}
	for _, factor := range factors {
		if i%factor.divisor == 0 {
			output += factor.result
		}
	}
	if output == "" {
		output = strconv.Itoa(i)
	}
	return output
}
