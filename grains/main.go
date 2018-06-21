package grains

import (
	"errors"
)

func Square(sq int) (uint64, error) {
	if sq < 1 || sq > 64 {
		return 0, errors.New("err")
	}
	var s uint64
	s = 1
	for i := 2; i <= sq; i++ {
		s *= 2
	}
	return s, nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		total, _ := Square(i)
		sum += total
	}

	return sum
}
