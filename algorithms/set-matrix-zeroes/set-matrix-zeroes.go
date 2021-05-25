// https://leetcode.com/problems/set-matrix-zeroes

package leetcode

// O(1) space
func setZeroes(matrix [][]int) {
	firstRow := false
	firstCol := false

	// check first row
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			firstRow = true
			break
		}
	}
	// check first column
	for i := range matrix {
		if matrix[i][0] == 0 {
			firstCol = true
			break
		}
	}

	// check except first row and first column
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// set zero except first row and first column
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// set zero for first row
	if firstRow {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}

	// set zero for first column
	if firstCol {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}

// O(m+n) space
func setZeroes2(matrix [][]int) {
	row := make([]int, len(matrix))
	col := make([]int, len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := range row {
		if row[i] == 0 {
			continue
		}
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}
	for i := range col {
		if col[i] == 0 {
			continue
		}
		for j := range matrix {
			matrix[j][i] = 0
		}
	}
}
