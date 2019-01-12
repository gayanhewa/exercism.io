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

func String2(phrase string) string {
	word := []rune(phrase)
	for i, j := 0, len(word)-1; i < len(word)/2; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return string(word)
}
