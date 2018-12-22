package strand

// ToRNA will convert DNA to RNA.
func ToRNA(dna string) string {

	var mapping = map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
	var output string
	for _, char := range dna {
		output += mapping[char]
	}

	return output
}
