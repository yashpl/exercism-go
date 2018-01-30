package grains

//Solution to calculate number of grains
var total_grains uint64

func Square(in int) (out uint64, err bool) {

	//initialize values
	out = 0
	err = false
	var previous_val uint64 = 1
	total_grains = 0

	//Check for valid paameters
	if in <= 0 || in > 64 {
		err = true
		return out, err
	}
	//if input is 1 , return 1
	if in < 2 {
		out = 1
		return out, err
	} else {
		for i := 2; i <= in; i++ {
      previous_val *= 2
      out = previous_val
			total_grains = total_grains + (previous_val)

		}
		out = total_grains
		return out, err
	}
}
func Total() uint64{
  return total_grains
}
