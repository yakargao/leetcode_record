package tree

var totalSum = 0
var treeNodeMap = make(map[int]int)

func pathSumIV(num []int) int {
	if len(num) == 0 {
		return 0
	}
	// 1. build a tree

	for _, n := range num {
		treeNodeMap[n/10] = n % 10
	}
	pathSumIVOfRoot(num[0]/10, 0)
	return totalSum
}

func pathSumIVOfRoot(pathPos int, sum int) {
	num, ok := treeNodeMap[pathPos]
	if !ok {
		return
	}
	sum += num
	left := (pathPos/10+1)*10 + ((pathPos%10)*2 - 1)
	right := left + 1
	_, ok1 := treeNodeMap[left]
	_, ok2 := treeNodeMap[right]
	if !ok1 && !ok2 {
		totalSum += sum
		return
	}
	pathSumIVOfRoot(left, sum)
	pathSumIVOfRoot(right, sum)
}
