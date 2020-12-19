package backtrack

var res1 []string

func permutation(s string) []string {
	res1 = make([]string, 0)
	permutationBT([]rune(s), make([]rune, 0), make([]bool, len(s)))
	return res1
}

func permutationBT(s, track []rune, check []bool) {
	if len(s) == len(track) {
		temp := make([]rune, len(track))
		copy(temp, track)
		res1 = append(res1, string(temp))
		return
	}
	for idx, v := range s {
		if check[idx] || idx > 0 && s[idx-1] == v && !check[idx-1] {
			continue
		}
		check[idx] = true
		track = append(track, v)
		permutationBT(s, track, check)
		check[idx] = false
		track = track[:len(track)-1]
	}
}
