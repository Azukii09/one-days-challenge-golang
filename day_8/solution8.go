package day_8

func IsSubset(first, second []Flight) bool {
	set := make(map[Flight]int)
	for _, flight2 := range second {
		set[flight2] += 1
	}
	for _, flight1 := range first {
		if count, found := set[flight1]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[flight1] = count - 1
		}
	}
	return true
}
