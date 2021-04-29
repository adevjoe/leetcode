// https://leetcode.com/problems/rotate-image

package leetcode

func rotate(matrix [][]int) {
	for i := 1; i <= len(matrix)/2; i++ {
		matrix[i-1], matrix[len(matrix)-i] = matrix[len(matrix)-i], matrix[i-1]
	}
	for i := range matrix {
		for j := range matrix[i] {
			if j-i > 0 {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
}
