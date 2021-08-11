// https://leetcode.com/problems/number-of-islands

package leetcode

func numIslands(grid [][]byte) int {
	land := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '1' {
				continue
			}
			// fill
			fill(grid, i, j)
			land++
		}
	}
	return land
}

func fill(grid [][]byte, i, j int) {
	grid[i][j] = 2
	// up
	if i > 0 && grid[i-1][j] == '1' {
		fill(grid, i-1, j)
	}
	// down
	if i < len(grid)-1 && grid[i+1][j] == '1' {
		fill(grid, i+1, j)
	}
	// left
	if j > 0 && grid[i][j-1] == '1' {
		fill(grid, i, j-1)
	}
	// right
	if j < len(grid[i])-1 && grid[i][j+1] == '1' {
		fill(grid, i, j+1)
	}
}
