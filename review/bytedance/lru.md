###题目描述
设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值
####[要求]
- set和get方法的时间复杂度为O(1)
- 某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
- 当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
对于每个操作2，输出一个答案
  
````
package main

import (
	"container/list"
	"fmt"
)

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	// write code here
	lru := NewLru(k)
	result := make([]int, 0)
	for idx, op := range operators {
		if len(op) == 0 {
			_ = fmt.Errorf("op len is illegal,%v", idx)
		}
		switch op[0] {
		case 1:
			//判断数组长度
			lru.set(op[1], op[2])
		case 2:
			result = append(result, lru.get(op[1]))
		}
	}
	return result
}

type Lru struct {
	ll     *list.List
	eleMap map[int]*list.Element
	limit  int
	size   int
}
type entry struct {
	key, value int
}

func NewLru(limit int) *Lru {
	return &Lru{
		ll:     list.New(),
		eleMap: make(map[int]*list.Element),
		limit:  limit,
		size:   0,
	}
}
func (lru *Lru) set(key, value int) {
	//如果本来就存在的，直接把元素移动到列表头部
	if ele, ok := lru.eleMap[key]; ok {
		lru.ll.MoveToFront(ele)
		ele.Value = &entry{key: key, value: value}
	} else { //如果不存在，直接加入
		if lru.size == lru.limit {
			lru.remove()
		}
		ele := lru.ll.PushFront(&entry{key: key, value: value})
		lru.eleMap[key] = ele
		lru.size++
	}
}
func (lru *Lru) remove() {
	ele := lru.ll.Back()
	if ele != nil {
		value := lru.ll.Remove(ele)
		entry := value.(*entry)
		delete(lru.eleMap, entry.key)
		lru.size--
	}
}
func (lru *Lru) get(key int) int {
	if ele, ok := lru.eleMap[key]; ok {
		lru.ll.MoveToFront(ele)
		entry := ele.Value.(*entry)
		return entry.value
	}
	return -1
}

````