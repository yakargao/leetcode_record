package tree

func verifyTreeOrder(postorder []int) bool {
	var isBst func(postorder []int, left, right int) bool

	isBst = func(postorder []int, left, right int) bool {
		if left >= right {
			return true
		}
		mid := left
		root := postorder[right]

		for postorder[mid] < root {
			mid++
		}
		for i := mid; i < right; i++ {
			if postorder[i] < root {
				return false
			}
		}
		return isBst(postorder, left, mid-1) && isBst(postorder, mid, right-1)
	}
	return isBst(postorder, 0, len(postorder)-1)

}
