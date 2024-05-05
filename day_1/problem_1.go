package day_1

import (
	"fmt"
	"strings"
)

/*
PairNumberForGivenSum
Problem day 1
Find a pair with the given sum in an array
Given an unsorted integer array, find a pair with the given sum in it.

For example,

Input: nums = [8, 7, 2, 5, 3, 1]
target = 10
Output: Pair found (8, 2)orPair found (7, 3)

Input: nums = [5, 2, 6, 8, 1, 9]
target = 12
Output: Pair not found
*/
func PairNumberForGivenSum(target int, nums []int) string {
	var output [][]int
	for index, n := range nums {
		for i := index + 1; i < len(nums); i++ {
			if index != i {
				if n+nums[i] == target {
					output = append(output, []int{n, nums[i]})
				}
			}
		}
	}
	if len(output) != 0 {
		var stringOutput string
		for indexFor, data := range output {
			if indexFor != len(output)-1 {
				stringOutput += "Pair found " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), ","), "[]") + " or "
			} else {
				stringOutput += "Pair found " + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data)), ","), "[]")
			}
		}
		return stringOutput
	}
	return "Pair not found"
}
