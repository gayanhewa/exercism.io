package raindrops

import (
	"fmt"
)

// Convert defines the method that converts numbers into raindrop speak.
func Convert(i int) string {
	var output string
	factor := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	for _, k := range [3]int{3, 5, 7} {
		if i%k == 0 {
			output += factor[k]
		}
	}
	if output == "" {
		output = fmt.Sprintf("%d", i)
	}
	return output
}
