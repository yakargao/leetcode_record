package sliding_window

import "math"

func minWindow(s string, t string) string {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range t {
		need[c] += 1
	}
	left, right := 0, 0
	valid := 0 //窗口中包含t字符串的个数
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := rune(s[right]) //移入窗口的字符
		right++             //移动
		//进行窗口内一系列数据的更新
		if _, ok := need[c]; ok { //加入的字符的t里面
			window[c]++
			if window[c] == need[c] { //达到要求
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := rune(s[left])
			//左移窗口
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
