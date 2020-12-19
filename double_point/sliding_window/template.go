package sliding_window

func slidingWidow(s, t string) {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c] += 1
	}
	left, right := 0, 0
	valid := 0 //窗口中包含t字符串的个数
	for right < len(s) {
		c := s[right] //移入窗口的字符
		right++       //移动
		//进行窗口内一系列数据的更新
		//TODO:应该更新哪些数据:window,valid
		window[rune(c)]++

		for valid == len(need) { //TODO：什么条件下应该缩小窗口：valid==len(need)
			d := s[left]
			//左移窗口
			left++
			//进行窗口内一系列数据的更新
			//TODO：应该更新哪些数据
			window[rune(d)] -= 1
		}
		//TODO：我们需要的结果应该在扩大窗口时还是缩小窗口时更新:缩小
	}

}
