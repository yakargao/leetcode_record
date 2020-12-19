package backtrack

import "fmt"

var res [][]int
func permute(nums []int) [][]int {
	res = make([][]int,0)
	track := make([]int,0)
	permuteBT(nums,track)
	return res
}


func permuteBT(nums []int,track []int)  {
	if len(nums) == len(track) { //可选列表为空
		fmt.Printf("%p|%v\n",&track,track)
		temp := make([]int, len(track))
		copy(temp,track)
		res = append(res,temp)
		return
	}
	for _,num := range  nums{
		//做选择
		if inSlice(num,track){
			continue
		}
		track = append(track,num)
		permuteBT(nums,track)
		//撤销选择
		track = track[:len(track)-1]
	}
}
func inSlice(num int,nums []int)bool  {
	for _,i := range nums{
		if num == i {
			return true
		}
	}
	return false
}
