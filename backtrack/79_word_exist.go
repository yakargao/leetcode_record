package backtrack

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右
var visited [][]bool

func exist(board [][]byte, word string) bool {
	rowLen, colLen := len(board), len(board[0])
	visited = make([][]bool, rowLen)
	for i := range visited {
		visited[i] = make([]bool, colLen)
	}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if backTrack(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false

}
func backTrack(board [][]byte, word string, i, j, k int) bool {
	if board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	visited[i][j] = true
	defer func() {
		visited[i][j] = false
	}()

	for _, dir := range directions {
		newI, newJ := i+dir.x, j+dir.y

		if newI >= 0 && newJ >= 0 && newI < len(board) && newJ < len(board[0]) {
			if visited[newI][newJ] {
				continue
			}

			if backTrack(board, word, newI, newJ, k+1) {
				return true
			}
		}
	}
	return false
}
