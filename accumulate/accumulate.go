package accumulate

// Accumulate defines the logic to accumulate the provided slice.
func Accumulate(input []string, f func(string) string) []string {
	for i, v := range input {
		input[i] = f(v)
	}
	return input
}
