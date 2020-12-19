package bfs

import (
	"container/list"
)

func openLock(deadends []string, target string) int {
	return openLockBFS2(deadends, target)
}
func upOne(curr []rune, idx int) []rune {
	newState := make([]rune, len(curr))
	copy(newState, curr)
	if curr[idx] == '9' {
		newState[idx] = '0'
	} else {
		newState[idx] += 1
	}
	return newState
}
func downOne(curr []rune, idx int) []rune {
	newState := make([]rune, len(curr))
	copy(newState, curr)
	if curr[idx] == '0' {
		newState[idx] = '9'
	} else {
		newState[idx] -= 1
	}
	return newState
}

func openLockBFS(deadends []string, target string) int {
	queue := list.New()
	visited := make(map[string]struct{})
	step := 0
	queue.PushBack([]rune("0000"))
	start := "0000"
	for _, d := range deadends { //将死胡同的步骤加到visitor
		visited[d] = struct{}{}
	}
	if _, ok := visited[start]; ok {
		return -1
	}
	visited[start] = struct{}{}
	for queue.Len() != 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			curr := queue.Front().Value.([]rune)
			queue.Remove(queue.Front())
			if string(curr) == target {
				return step
			}
			for idx := 0; idx < 4; idx++ {
				up := upOne(curr, idx)
				if _, ok := visited[string(up)]; !ok {
					queue.PushBack(up)
					visited[string(up)] = struct{}{}
				}
				down := downOne(curr, idx)
				if _, ok := visited[string(down)]; !ok {
					queue.PushBack(down)
					visited[string(down)] = struct{}{}
				}
			}
		}
		step++
	}
	return -1
}

//双向优化
func openLockBFS2(deadends []string, target string) int {
	visited := make(map[string]struct{})
	set1 := make(map[string]struct{})
	set2 := make(map[string]struct{})
	start := "0000"
	for _, d := range deadends { //将死胡同的步骤加到visitor
		visited[d] = struct{}{}
	}
	if _, ok := visited[start]; ok {
		return -1
	}
	if _, ok := visited[target]; ok {
		return -1
	}
	step := 0
	set1[start] = struct{}{}
	set2[target] = struct{}{}
	for len(set1) != 0 && len(set2) != 0 {
		temp := make(map[string]struct{})
		for curr, _ := range set1 {
			if _, ok := visited[curr]; ok {
				continue
			}
			if _, ok := set2[curr]; ok {
				return step
			}
			visited[curr] = struct{}{}
			for idx := 0; idx < 4; idx++ {
				up := upOne([]rune(curr), idx)
				if _, ok := visited[string(up)]; !ok {
					temp[string(up)] = struct{}{}
				}
				down := downOne([]rune(curr), idx)
				if _, ok := visited[string(down)]; !ok {
					temp[string(down)] = struct{}{}
				}
			}
		}
		step++
		set1 = set2
		set2 = temp
	}
	return -1
}
