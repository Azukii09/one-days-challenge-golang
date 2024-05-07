package day_1

import (
	"fmt"
	"strings"
)

func PairNumberForGivenSum(target int, nums []int) string {
	var output [][]int
	//loop for get output of pair
	for index, n := range nums {
		for i := index + 1; i < len(nums); i++ {
			if index != i {
				if n+nums[i] == target {
					output = append(output, []int{n, nums[i]})
				}
			}
		}
	}

	//print the output
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
