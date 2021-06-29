// https://leetcode.com/problems/pascals-triangle

package leetcode

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{1})
	for i := 2; i <= numRows; i++ {
		row := make([]int, 0)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				row = append(row, 1)
				continue
			}
			row = append(row, result[i-2][j-1]+result[i-2][j])

		}
		result = append(result, row)
	}
	return result
}
