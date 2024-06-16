package doublePoint

func lengthOfLongestSubstring(s string) int {
	charCounter := make(map[rune]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		charCounter[c]++
		for charCounter[c] > 1 {
			d := rune(s[left])
			left++
			charCounter[d]--
		}
		res = max(res, right-left)
	}
	return res
}
