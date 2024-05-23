package day_19

func getDigit(chr byte) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

func ISBN(input string) bool {
	length := len(input)
	pos := 10
	sum := 0

	for i := 0; i < length; i++ {
		chr := input[i]
		if digit, ok := getDigit(chr); ok {
			sum += (digit * pos)
			pos--
		} else if pos == 1 && chr == 'X' {
			sum += 10
			pos--
		} else if chr != '-' {
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}
