package graph

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}
func dfs(grid [][]byte, r, c int) {
	// 判断 base case
	// 如果坐标 (r, c) 超出了网格范围，直接返回
	if !inArea(grid, r, c) {
		return
	}
	if grid[r][c] != '1' {
		return
	}
	grid[r][c] = '2' // 将格子标记为「已遍历过」
	for _, dir := range directions {
		dfs(grid, r+dir[0], c+dir[1])
	}
	//// 访问上、下、左、右四个相邻结点
	//dfs(grid, r-1, c)
	//dfs(grid, r+1, c)
	//dfs(grid, r, c-1)
	//dfs(grid, r, c+1)
}

// 判断坐标 (r, c) 是否在网格中
func inArea(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
