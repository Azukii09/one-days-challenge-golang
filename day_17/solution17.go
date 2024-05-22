package day_17

func NumOfStep(num int) int {
	step := 0

	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else if num%2 == 1 {
			num = (num * 3) + 1
		}
		step++
	}
	return step
}
