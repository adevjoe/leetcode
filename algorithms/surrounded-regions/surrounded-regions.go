// https://leetcode.com/problems/surrounded-regions

package leetcode

func solve(board [][]byte) {
	// find 'O' from line 1 and last line
	for i := 0; i < len(board)-1; i++ {
		if board[i][0] == 'O' {
			fill(board, i, 0)
		}
		if board[i][len(board[i])-1] == 'O' {
			fill(board, i, len(board[i])-1)
		}
	}
	// find 'O' from row 1 and last row
	for i := 0; i < len(board[0])-1; i++ {
		if board[0][i] == 'O' {
			fill(board, 0, i)
		}
		if board[len(board)-1][i] == 'O' {
			fill(board, len(board)-1, i)
		}
	}

	// fill captured region
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i > 0 && j > 0 && i < len(board)-1 && j < len(board[i])-1 && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 0 {
				board[i][j] = 'O'
			}
		}
	}
}

// board[i][j] 0 => not captured
func fill(board [][]byte, i, j int) {
	board[i][j] = 0
	// up
	if i > 0 && board[i-1][j] == 'O' {
		fill(board, i-1, j)
	}
	// down
	if i < len(board)-1 && board[i+1][j] == 'O' {
		fill(board, i+1, j)
	}
	// left
	if j > 0 && board[i][j-1] == 'O' {
		fill(board, i, j-1)
	}
	// right
	if j < len(board[i])-1 && board[i][j+1] == 'O' {
		fill(board, i, j+1)
	}
}
