// Package twofer implements functions needed for two for one.
package twofer

import (
	"fmt"
)

// ShareWith provides the two for one functionality.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
