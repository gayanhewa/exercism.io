package strand

// ToRNA will convert DNA to RNA.
func ToRNA(dna string) string {

	var mapping = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	output := make([]rune, len(dna))
	for i, char := range dna {
		output[i] = mapping[char]
	}
	return string(output)
}
