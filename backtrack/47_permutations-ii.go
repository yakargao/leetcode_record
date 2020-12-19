package backtrack

import (
	"sort"
)


func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res = make([][]int,0)
	track := make([]int,0)
	check := make([]bool,len(nums))
	permuteUniqueTB(nums,track,check)
	return res
}
func permuteUniqueTB(nums,track []int,check []bool)  {
	if len(nums) == len(track) {
		tmp := make([]int,len(track))
		copy(tmp,track)
		res = append(res,tmp)
		return
	}
	for idx,num := range nums {
		if check[idx] ||  idx > 0 && nums[idx] == nums[idx-1]&&!check[idx-1]  { //剪枝
			continue
		}
		check[idx] = true
		track = append(track,num)
		permuteUniqueTB(nums,track,check)
		track = track[:len(track)-1]
		check[idx] = false
	}
}