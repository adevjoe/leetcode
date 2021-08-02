// https://leetcode.com/problems/excel-sheet-column-number

package leetcode

import "math"

func titleToNumber(columnTitle string) int {
	l := len(columnTitle)
	result := 0
	for i := 0; i < l; i++ {
		result += int(columnTitle[i]-'A'+1) * int(math.Pow(26, float64(l-i-1)))
	}
	return result
}
