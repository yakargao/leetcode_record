package matrix

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	if len(matrix) == 0 {
		return ans
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for {
		//从左到右
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		//从上到下
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		//从右到左
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		//从下到上
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
