package day_21

// Return all substrings of size n within s
func All(n int, s string) []string {
	// If asking for more than provided, get nothing
	if n > len(s) {
		return nil
	}
	subStrings := make([]string, 0, len(s)-n)
	for i := 0; i+n <= len(s); i++ {
		subStrings = append(subStrings, s[i:i+n])
	}
	return subStrings
}

func UnsafeFirst(n int, s string) string {
	subStrings := All(n, s)
	if subStrings == nil {
		return ""
	}
	return subStrings[0]
}
