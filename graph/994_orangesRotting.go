package graph

type Point struct {
	x, y int
}

var dirs = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := make([]Point, 0)
	fresh := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, Point{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(queue) > 0 {
		minutes++
		size := len(queue)
		for i := 0; i < size; i++ {
			point := queue[0]
			queue = queue[1:]
			for _, dir := range dirs {
				x, y := point.x+dir.x, point.y+dir.y
				if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				fresh--
				queue = append(queue, Point{x, y})
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return minutes - 1
}
