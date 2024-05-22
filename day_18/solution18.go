package day_18

import (
	"strconv"
	"strings"
)

func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	if _, err := strconv.Atoi(num); err != nil {
		return false
	}
	total := 0
	pos := 0
	for i := len(num) - 1; i > -1; i-- {
		digit := int(num[i] - '0')
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
