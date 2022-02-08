package series

func All(n int, s string) []string {
	substrings := []string{}
	for i := 0; i <= len(s)-n; i++ {
		substrings = append(substrings, s[i:i+n])
	}
	return substrings
}

func UnsafeFirst(n int, s string) string {
	if n <= len(s) {
		return s[0:n]
	}
	return ""
}
