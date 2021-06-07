// https://leetcode.com/problems/word-search

package leetcode

var cache [][]bool

func exist(board [][]byte, word string) bool {
	cache = make([][]bool, len(board))
	for m := range board {
		cache[m] = make([]bool, len(board[m]))
	}
	for m := range board {
		for n := range board[m] {
			if board[m][n] == word[0] {
				cache[m][n] = true
				if len(word) == 1 {
					return true
				}
				if existPlus(board, word[1:], m, n) {
					return true
				}
			}
		}
	}
	return false
}

func existPlus(board [][]byte, word string, m, n int) bool {
	if word == "" {
		return true
	}
	// check top
	if m != 0 && !cache[m-1][n] {
		if board[m-1][n] == word[0] {
			cache[m-1][n] = true
			if existPlus(board, word[1:], m-1, n) {
				return true
			}
		}
	}

	// check bottom
	if m != len(board)-1 && !cache[m+1][n] {
		if board[m+1][n] == word[0] {
			cache[m+1][n] = true
			if existPlus(board, word[1:], m+1, n) {
				return true
			}
		}
	}

	// check left
	if n != 0 && !cache[m][n-1] {
		if board[m][n-1] == word[0] {
			cache[m][n-1] = true
			if existPlus(board, word[1:], m, n-1) {
				return true
			}
		}
	}

	// check right
	if n != len(board[m])-1 && !cache[m][n+1] {
		if board[m][n+1] == word[0] {
			cache[m][n+1] = true
			if existPlus(board, word[1:], m, n+1) {
				return true
			}
		}
	}
	cache[m][n] = false
	return false
}
