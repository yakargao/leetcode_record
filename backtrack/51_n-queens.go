package backtrack

var nqRes [][]string

func solveNQueens(n int) [][]string {
	src := make([][]rune, n)
	for i := 0; i < n; i++ {
		row := make([]rune, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		src[i] = row
	}

	nqRes = make([][]string, 0)
	solveNQueensBT(src, 0)
	return nqRes
}
func solveNQueensBT(src [][]rune, row int) {
	l := len(src)
	if row == l {
		r := make([]string, l)
		for i := 0; i < l; i++ {
			r[i] = string(src[i])
		}
		nqRes = append(nqRes, r)
		return
	}
	for col := 0; col < len(src); col++ {
		if !isValid(src, row, col) {
			continue
		}
		src[row][col] = 'Q'
		solveNQueensBT(src, row+1)
		src[row][col] = '.'

	}
}
func isValid(src [][]rune, row, col int) bool {
	//检查同列
	for i := 0; i < row; i++ {
		if src[i][col] == 'Q' {
			return false
		}
	}
	//检查左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if src[i][j] == 'Q' {
			return false
		}
	}
	//检查右上
	for i, j := row-1, col+1; i >= 0 && j < len(src); i, j = i-1, j+1 {
		if src[i][j] == 'Q' {
			return false
		}
	}
	return true
}
