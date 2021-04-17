package cn

import "container/list"

//栈排序。 编写程序，对栈进行排序使最小元素位于栈顶。最多只能使用一个其他的临时栈存放数据，但不得将元素复制到别的数据结构（如数组）中。该栈支持如下操作：pu
//sh、pop、peek 和 isEmpty。当栈为空时，peek 返回 -1。 
//
// 示例1: 
//
//  输入：
//["SortedStack", "push", "push", "peek", "pop", "peek"]
//[[], [1], [2], [], [], []]
// 输出：
//[null,null,null,1,null,2]
// 
//
// 示例2: 
//
//  输入： 
//["SortedStack", "pop", "pop", "push", "pop", "isEmpty"]
//[[], [], [], [1], [], []]
// 输出：
//[null,null,null,null,null,true]
// 
//
// 说明: 
//
// 
// 栈中的元素数目在[0, 5000]范围内。 
// 
// Related Topics 设计 
// 👍 33 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type SortedStack struct {
	stack *list.List
	help *list.List
}


func Constructor() SortedStack {
	return SortedStack{
		stack: list.New(),
		help: list.New(),
	}
}


func (this *SortedStack) Push(val int)  {
	if this.stack.Len() == 0 {
		this.stack.PushBack(val)
	}else{
		for this.stack.Len() != 0&&
			this.stack.Back().Value.(int) < val{
			this.help.PushBack(this.stack.Back().Value.(int))
			this.stack.Remove(this.stack.Back())
		}
		this.stack.PushBack(val)
		for this.help.Len()!=0{
			this.stack.PushBack(this.help.Remove(this.help.Back()))
		}
	}
}


func (this *SortedStack) Pop()  {
	if this.stack.Len() != 0 {
		this.stack.Remove(this.stack.Back())
	}
}


func (this *SortedStack) Peek() int {
	if this.stack.Len() != 0 {
		return this.stack.Back().Value.(int)
	}
	return  -1
}


func (this *SortedStack) IsEmpty() bool {
	return this.stack.Len() ==0
}


/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
//leetcode submit region end(Prohibit modification and deletion)
