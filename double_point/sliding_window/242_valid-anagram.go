package sliding_window

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	m := make(map[rune]int)
	for _, i := range s {
		m[i]++
	}
	for _, i := range t {
		m[i]--
	}
	for _, i := range s {
		if m[i] > 0 {
			return false
		}
	}
	return true
}
