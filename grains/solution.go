package grains

import (
	"errors"
)

//Solution to calculate number of grains
var total_grains uint64

func Square(in int) (uint64, error) {

	//initialize values
	var out uint64 = 0
	err := errors.New("input out of range")
	var previous_val uint64 = 1
	total_grains = 0

	//Check for valid paameters
	if in <= 0 || in > 64 {

		return out, err
	}
	//if input is 1 , return 1 ( Used HACK below )
	if in < 2 {
		out = 1
		return out, nil
	} else {
		for i := 2; i <= in; i++ {
			previous_val *= 2
			out = previous_val
			total_grains = total_grains + (previous_val)

		}

		return out, nil
	}
}
func Total() uint64 {
	Square(64)
	// HACK: one is added to return value to compensate condition in line 24 ( if in < 2 )
	return total_grains + 1
}
