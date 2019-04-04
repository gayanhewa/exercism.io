package diffsquares

// SquareOfSum returns the square of sum for n natural numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares returns the square of sum for n natural numbers.
func SumOfSquares(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference between SquareOfSum and SumOfSquares.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
