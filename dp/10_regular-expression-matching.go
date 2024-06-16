package dp

import "fmt"

var memo map[string]bool

func isMatch(s string, p string) bool {
	memo = make(map[string]bool)
	return isMatchDP(s, 0, p, 0)
}

func isMatchDP(s string, i int, p string, j int) bool {

	len1, len2 := len(s), len(p)
	if j == len2 {
		return i == len1
	}

	if i == len1 {
		if (len2-j)%2 == 1 {
			return false
		}
		for ; j+1 < len2; j += 2 {
			if rune(p[j+1]) != '*' {
				return false
			}
		}
		return true
	}
	res := false
	memoKey := fmt.Sprintf("%v#%v", i, j)
	if v, ok := memo[memoKey]; ok {
		return v
	}
	if s[i] == p[j] || p[j] == '.' {
		if j < len2-1 && p[j+1] == '*' {
			res = isMatchDP(s, i, p, j+2) || isMatchDP(s, i+1, p, j)
		} else {
			res = isMatchDP(s, i+1, p, j+1)
		}

	} else {
		if j < len2-1 && p[j+1] == '*' {
			res = isMatchDP(s, i, p, j+2)
		} else {
			res = false
		}
	}
	memo[memoKey] = res
	return res
}
