package day_12

import (
	"errors"
	"sort"
)

var romanNumeralDict = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func RomanNumbers(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New("Invalid")
	}
	keys := []int{}
	for k, _ := range romanNumeralDict {
		keys = append(keys, k)
	}
	// Sort map keys in reverse order to formulate the roman numeral
	// by subtracting highest to lowest
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	roman := ""
	for _, k := range keys {
		for input >= k {
			roman += romanNumeralDict[k]
			input += -k
		}
	}
	return roman, nil
}
