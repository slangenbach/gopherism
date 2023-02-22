package grains

import (
	"fmt"
	"math"
)

// Calcuate the amount of grain on a square given its number on the chessboard
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, fmt.Errorf("%d is not a valid square", number)
	} else if number == 1 {
		return 1, nil
	} else {
		square, _ := Square(number - 1)
		return 2 * square, nil
	}
}

// "Calculate" the total amount of grain on a chessboard with 64 squares
func Total() uint64 {
	return math.MaxUint64
}
