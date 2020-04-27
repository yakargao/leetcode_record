/**
* @Author: CiachoG
* @Description：
* @Date: 2020/4/21 10:52
 */
package algorithm

import (
	"fmt"
	"testing"
)

func TestWordTree(t *testing.T) {
	wordTree := new(WordDictionary)
	wordTree.AddWord("bad")
	wordTree.AddWord("dad")
	wordTree.AddWord("mad")
	fmt.Println(wordTree.Search("pad"))
	fmt.Println(wordTree.Search("bad"))

}
