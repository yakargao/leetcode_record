package algorithm

import (
	"fmt"
	"strconv"
	"testing"
)

var ops = []string{"Trie", "insert", "insert", "insert", "insert", "insert", "insert", "search", "search", "search", "search", "search", "search", "search", "search", "search", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith", "startsWith"}
var values = []string{"", "app", "apple", "beer", "add", "jam", "rental", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam", "apps", "app", "ad", "applepie", "rest", "jan", "rent", "beer", "jam"}
var answers = []string{"null", "null", "null", "null", "null", "null", "null", "false", "true", "false", "false", "false", "false", "false", "true", "true", "false", "true", "true", "false", "false", "false", "true", "true", "true"}

func Test105(t *testing.T) {
	trie := new(Trie)
	for idx, op := range ops {
		switch op {
		case "Trie":
			fmt.Println("null")
		case "insert":
			trie.Insert(values[idx])
			fmt.Print("null ")
		case "search":
			result := trie.Search(values[idx])
			if strconv.FormatBool(result) != answers[idx] {
				fmt.Print("search", values[idx])
			}
		case "startsWith":
			result := trie.StartsWith(values[idx])
			if strconv.FormatBool(result) != answers[idx] {
				fmt.Print("StartsWith", values[idx])
			}
		}
	}
	fmt.Println(trie)
}

//[null,null,null,null,null,null,null,false,true,false,false,false,false,false,true,true,false,true,true,false,false,false,true,true,true]
