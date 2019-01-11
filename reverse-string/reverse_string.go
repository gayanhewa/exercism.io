package reverse

// String returns a reversed string.
func String(phrase string) string {
	word := []rune(phrase)
	var reversed []rune
	for i := len(word) - 1; i >= 0; i-- {
		reversed = append(reversed, word[i])
	}
	return string(reversed)
}
