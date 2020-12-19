package sliding_window

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	need, window := make(map[rune]int), make(map[rune]int)
	for _, n := range p {
		need[n]++
	}
	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		//考虑left右移动
		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			b := rune(s[left])
			left++
			if _, ok := need[b]; ok {
				if window[b] == need[b] {
					valid--
				}
				window[b]--
			}
		}
	}
	return res
}
