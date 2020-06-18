/**
* @Author: CiachoG
* @Description：
* @Date: 2020/4/20 21:34
 */
package algorithm

type WordDictionary struct {
	Val      string
	Children []*WordDictionary
	End      bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{End: false}
}

func (this *WordDictionary) matchChild(val string) *WordDictionary {
	for _, child := range this.Children {
		if val == child.Val {
			return child
		}
	}
	return nil
}

func (this *WordDictionary) addWord(word string, height int) {
	if height == len(word) {
		this.End = true
		return
	}
	val := word[height]
	child := this.matchChild(string(val))
	if child == nil {
		child = &WordDictionary{Val: string(val), End: false}
		this.Children = append(this.Children, child)
	}
	child.addWord(word, height+1)
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.addWord(word, 0)
}

func (this *WordDictionary) search(word string, height int) bool {
	if len(word) == height {
		val := string(word[height-1])
		if this.Val == val && this.End || val == "." && this.End {
			return true
		}
		return false
	}
	val := string(word[height])
	for _, child := range this.Children {
		if child.Val == val || val == "." {
			result := child.search(word, height+1)
			if result == true {
				return result
			}
		}
	}
	return false
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.search(word, 0)
}
