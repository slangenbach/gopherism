package diffsquares

import "math"

// Calculate square of sum of first n natural numbers
func SquareOfSum(n int) int {
	var sum int

	for i := n; i > 0; i-- {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// Calculate sum of squares of first n natural numbers
func SumOfSquares(n int) int {
	var sum int

	for i := n; i > 0; i-- {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
