package day_2

import (
	"math"
)

func LCS(str1 string, str2 string, finalValue int) int {
	//get string length
	str1len := len(str1)
	str2len := len(str2)

	//make last char variable
	var car1 string
	var car2 string

	// get last character in the string
	if str1len > 0 && str2len > 0 {
		car1 = string(str1[str1len-1:])
		car2 = string(str2[str2len-1:])
	}

	// first check for the last character in each string is same or not
	if car1 == car2 && (str1len > 0 && str2len > 0) {
		//take out the last character in the string
		str1 = string(str1[0 : str1len-1])
		str2 = string(str2[0 : str2len-1])

		// if same add 1 to finalValue variable and start recursive function. then return the finalValue
		finalValue = 1 + LCS(str1, str2, finalValue)
		return finalValue

		// second check for the last character in each string is same or not
	} else if car1 != car2 && (str1len > 0 && str2len > 0) {

		/* if not same calculate the max value of finalValue in this condition:
		   1. we remove last character in string2, and we do start recursive function.
		   2. we remove last character in string1, and we do start recursive function.
		   3. check the max value from finalValue between two-step above
		*/
		finalValue = int(math.Max(float64(LCS(string(str1[0:str1len-1]), str2, finalValue)), float64(LCS(str1, string(str2[0:str2len-1]), finalValue))))
	} else {
		finalValue += 0
	}
	return finalValue
}
