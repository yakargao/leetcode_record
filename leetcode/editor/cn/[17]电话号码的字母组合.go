package cn

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 
//
// 示例 2： 
//
// 
//输入：digits = ""
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：digits = "2"
//输出：["a","b","c"]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= digits.length <= 4 
// digits[i] 是范围 ['2', '9'] 的一个数字。 
// 
// Related Topics 深度优先搜索 递归 字符串 回溯算法 
// 👍 1250 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
var res []string
var  letterMap = map[rune]string{'2':"abc",'3':"def",'4':"ghi",'5':"jkl",'6':"mno",'7':"pqrs",'8':"tuv",'9':"wxyz"}
func letterCombinations(digits string) []string {
	res = make([]string,0)
	if len(digits) == 0 {
		return res
	}
	backTrack(digits,[]rune{},len(digits),0)
	return res
}
func backTrack(digits string,track []rune,l,idx int)  {
	if len(track) == l {
		temp := make([]rune,len(track))
		copy(temp,track)
		res = append(res,string(temp))
		//fmt.Println(res)
		return
	}
	for i:=idx;i<len(digits);i++{
		for _,letter := range letterMap[rune(digits[i])]{
			track = append(track,letter)
			backTrack(digits,track,l,i+1)
			track = track[:len(track)-1]
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
