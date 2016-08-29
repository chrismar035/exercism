package grains

import (
	"errors"
	"math"
)

// Square computes the amount of grain on the given square
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("This number is not on the board")
	}

	return uint64(math.Pow(2, float64(num-1))), nil
}

// Total returns the total number of grains on the board
func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		num, _ := Square(i)
		sum += num
	}

	return sum
}
