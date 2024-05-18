package day_14

import "math"

func IsNumber(n int) bool {
	temp := n
	numDigits := 0
	for temp > 0 {
		temp /= 10
		numDigits++
	}
	// Sum the digits raised to the power of numDigits
	sum := 0
	temp = n
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}
	// Check if the sum is equal to the original number
	return sum == n
}
