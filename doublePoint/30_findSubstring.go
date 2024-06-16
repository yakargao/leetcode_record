package doublePoint

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	if len(s) == 1 && len(words) == 1 && words[0] == s {
		return []int{0}
	}

	wordLen := len(words[0])
	wordNum := len(words)
	left, right := 0, 0
	wordCounter := make(map[string]int)
	wordNeed := make(map[string]int)
	for _, word := range words {
		wordNeed[word]++
	}
	validNum := 0
	res := make([]int, 0)
	for right+wordLen <= len(s) {
		rangeWord := s[right : right+wordLen]
		right += 1

		if _, ok := wordNeed[rangeWord]; ok {
			wordCounter[rangeWord]++
			if wordCounter[rangeWord] <= wordNeed[rangeWord] {
				validNum++
			}
		}

		for right-left >= wordLen*wordNum {
			if validNum == wordNum {
				res = append(res, left)
			}

			rangeWord := s[left : left+wordLen]
			left += 1

			if _, ok := wordNeed[rangeWord]; ok {
				if wordCounter[rangeWord] <= wordNeed[rangeWord] {
					validNum--
				}
				wordCounter[rangeWord]--
			}
		}
	}
	return res
}
