package luhn

import (
	"strconv"
	"strings"
)

//Validate the card number
func Valid(in string) bool {
	is2nd := false
	total := 0
  in = strings.Replace(in," ","",-1)
	//Iterate from right to left
	for i := len(in) - 1; i >= 0; i-- {
		//Check if current character is a number
		if strings.ContainsAny(string(in[i]), "1234567890 ()") {
			//convert string to integer
			if strings.ContainsAny(string(in[i]), "1234567890") {
				ConvNum, _ := strconv.Atoi(string(in[i]))

				//check if current number is "every 2nd from right"
				if is2nd {
					//check if num*2 is greater than 9
					if (ConvNum * 2) >= 9 {
						ConvNum = ConvNum * 2
						ConvNum = ConvNum - 9
					} else {
						//doubling number is less than 9
						ConvNum = ConvNum * 2
					}
					total = total + ConvNum
					is2nd = false
				} else {
					//add other number directly
					total = total + ConvNum
					is2nd = true

				}
			} else { // not a number
				continue
			}
		} else {
			//False if anything other than int comes
			return false

		}
	}
	// check if total is divisible by 10
	if total%10 == 0 && len(in) > 1 {
		return true
	} else {
		return false
	}
}
