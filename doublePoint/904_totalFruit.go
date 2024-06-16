package doublePoint

func totalFruit(fruits []int) int {
	left, right := 0, 0
	res := 0
	count := 0
	categoryMap := make(map[int]int)
	for right < len(fruits) {
		categoryMap[fruits[right]]++
		count++
		right++
		for len(categoryMap) > 2 {
			categoryMap[fruits[left]]--
			if categoryMap[fruits[left]] == 0 {
				delete(categoryMap, fruits[left])
			}
			count--
			left++
		}
		res = max(res, count)
	}
	return res
}
