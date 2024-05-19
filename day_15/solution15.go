package day_15

import "strconv"

func Raindrop(number int) string {
	replacement := ""
	pling := number%3 == 0
	plang := number%5 == 0
	plong := number%7 == 0
	if pling {
		replacement += "Pling"
	}
	if plang {
		replacement += "Plang"
	}
	if plong {
		replacement += "Plong"
	}
	if !(pling || plang || plong) {
		return strconv.Itoa(number)
	}
	return replacement
}
