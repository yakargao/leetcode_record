/**
* @Author: CiachoG
* @Description：
* @Date: 2020/4/21 17:46
 */
package algorithm

import (
	"fmt"
	"testing"
)

func Test211(t *testing.T) {
	trie := Constructor2()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 true

}
