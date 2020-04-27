/**
* @Author: CiachoG
* @Description：
* @Date: 2020/4/21 17:34
 */
package algorithm

import (
	"strings"
)

type Trie2 struct {
	wordMap map[int][]string
	maxLen  int
}

/** Initialize your data structure here. */
func Constructor2() Trie2 {
	return Trie2{wordMap: make(map[int][]string), maxLen: 0}
}

/** Inserts a word into the trie. */
func (this *Trie2) Insert(word string) {
	s := this.wordMap[len(word)]
	s = append(s, word)
	if len(word) > this.maxLen {
		this.maxLen = len(word)
	}
	this.wordMap[len(word)] = s
}

/** Returns if the word is in the trie. */
func (this *Trie2) Search(word string) bool {
	if val, ok := this.wordMap[len(word)]; !ok {
		return false
	} else {
		for _, item := range val {
			if item == word {
				return true
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie2) StartsWith(prefix string) bool {
	preLen := len(prefix)
	for i := preLen; i <= this.maxLen; i++ {
		if val, ok := this.wordMap[i]; ok {
			for _, item := range val {
				if strings.HasPrefix(item, prefix) {
					return true
				}
			}
		}
	}
	return false
}
