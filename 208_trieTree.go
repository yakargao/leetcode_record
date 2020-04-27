/**
* @Author: CiachoG
* @Description：
* @Date: 2020/4/18 14:45
 */
package algorithm

type Trie struct {
	val      string
	children []*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func constructor() Trie {
	return Trie{}
}

//匹配某个节点,用于插入
func (t *Trie) matchChild(val string) *Trie {
	for _, child := range t.children {
		if child.val == val {
			return child
		}
	}
	return nil
}
func (t *Trie) insert(word string, height int) {
	if len(word) == height {
		t.isEnd = true
		return
	}
	val := word[height]
	child := t.matchChild(string(val))
	if child == nil {
		child = &Trie{val: string(val), isEnd: false}
		t.children = append(t.children, child)
	}
	child.insert(word, height+1)

}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	t.insert(word, 0)
}

func (t *Trie) search(word string, height int) bool {
	if len(word) == height {
		if t.val != string(word[height-1]) || !t.isEnd {
			return false
		}
		return true
	}
	for _, child := range t.children {
		if child.val == string(word[height]) {
			result := child.search(word, height+1)
			if result == true {
				return result
			}
		}
	}
	return false
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	return t.search(word, 0)
}

func (t *Trie) startWith(prefix string, height int) bool {
	if height >= len(prefix) {
		if t.val != string(prefix[height-1]) {
			return false
		}
		return true
	}
	for _, child := range t.children {
		if child.val == string(prefix[height]) {
			result := child.startWith(prefix, height+1)
			if result == true {
				return result
			}
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.startWith(prefix, 0)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
