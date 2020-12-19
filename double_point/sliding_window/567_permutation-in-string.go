package sliding_window

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, s := range s1 {
		need[s] += 1
	}
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		//right滑动
		c := rune(s2[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := rune(s2[left])
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	return false
}
