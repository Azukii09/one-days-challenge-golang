package day_4

import (
	"strings"
)

func PermutationCheck(str1, str2 string) bool {
	str1len := len(str1)
	str2len := len(str2)

	var count int

	if str1len != str2len {
		return false
	}

	for i := 0; i < str1len; i++ {
		if strings.Contains(str2, str1[i:i+1]) {
			count++
		}
	}
	if count == str1len {
		return true
	} else {
		return false
	}
}
