package day_9

import (
	"strconv"
	"strings"
)

var resistorColor = map[string]int{
	"Black":  0,
	"Brown":  1,
	"Red":    2,
	"Orange": 3,
	"Yellow": 4,
	"Green":  5,
	"Blue":   6,
	"Violet": 7,
	"Grey":   8,
	"White":  9,
}

func ResistorColorConverter(input []string) string {
	convertedResistor := ""
	var str string
	for i, s := range input {
		str = strings.ToUpper(s[0:1]) + s[1:len(s)]
		number := resistorColor[str]
		if i != 2 {
			convertedResistor += strconv.Itoa(number)
		}
	}
	if resistorColor[str] >= 3 && resistorColor[str] < 6 {
		for y := 0; y < (resistorColor[str] - 3); y++ {
			convertedResistor += "0"
		}
		convertedResistor += " kilo"
	} else if resistorColor[str] >= 6 && resistorColor[str] < 9 {
		for y := 0; y < (resistorColor[str] - 6); y++ {
			convertedResistor += "0"
		}
		convertedResistor += " mega"
	} else if resistorColor[str] >= 9 {
		for y := 0; y < (resistorColor[str] - 9); y++ {
			convertedResistor += "0"
		}
		convertedResistor += " giga"
	} else {
		for y := 0; y < resistorColor[str]; y++ {
			convertedResistor += "0"
		}
	}
	convertedResistor += " ohm"
	return convertedResistor
}
