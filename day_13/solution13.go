package day_13

import (
	"slices"
)

func MultipleSum(lvl int, mgcNum []int) int {
	var data []int
	var finalSum int
	for _, num := range mgcNum {
		for i := 1; i < lvl; i++ {
			if i%num == 0 {
				if !slices.Contains(data, i) {
					data = append(data, i)
				}
			}
		}
	}
	for _, num := range data {
		finalSum += num
	}
	return finalSum
}
