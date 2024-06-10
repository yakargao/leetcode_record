package binarySearch

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		}
	}
	return false

}
