package day_28

import (
	"errors"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}
	if n == 1 {
		return 2, nil
	}
	primeCount := 1
	for i := 3; ; i += 2 {
		if isPrime(i) {
			primeCount++
			if primeCount == n {
				return i, nil
			}
		}
	}
}
