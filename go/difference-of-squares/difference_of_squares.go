package diffsquares

// formulas ref: https://brilliant.org/wiki/sum-of-n-n2-or-n3/

// SquareOfSum returns the square of sum of first n natural numbers
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares returns the sum of squares of first n natural numbers
func SumOfSquares(n int) int {
	return ((n * (n + 1)) * (2*n + 1)) / 6
}

// Difference returns the difference of square of sum and sum of square of n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
