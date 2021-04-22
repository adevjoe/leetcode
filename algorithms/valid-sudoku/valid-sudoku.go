// https://leetcode.com/problems/valid-sudoku

package leetcode

import "fmt"

/**
规则大意是每行、每列、每子方块(3x3)中的数字是唯一的。
思路：
暴力解法就是每行、每列、每子方块单独检测一次，先遍历吧行、列、子方块弄成 3*9 个数组，
然后依次判断这 27 个数组是否重复，判断重复可以先排序，然后遍历一遍。时间复杂度为 O(n^2) + O(logn) + O(n)，空间复杂度为 n^3

更优的解法是全部遍历一遍，遍历时用 hashmap 保存已经每行、每列、每子方块已有的数字，如：
第一行第一个数字 1 可以计为 r1-1，列可以计为 c1-1，块可以计为 b1-1-1，遍历是判断该数字是否已经存在即可。
时间复杂度为 O(n^2)，空间复杂度为 O(n^2)
*/
func isValidSudoku(board [][]byte) bool {
	m := make(map[string]interface{})
	for r, b := range board {
		for c, item := range b {
			if item == '.' {
				continue
			}
			if _, ok := m[fmt.Sprintf("r%d-%d", r, item)]; ok {
				return false
			} else {
				m[fmt.Sprintf("r%d-%d", r, item)] = struct{}{}
			}
			if _, ok := m[fmt.Sprintf("c%d-%d", c, item)]; ok {
				return false
			} else {
				m[fmt.Sprintf("c%d-%d", c, item)] = struct{}{}
			}
			if _, ok := m[fmt.Sprintf("b%d-%d-%d", r/3, c/3, item)]; ok {
				return false
			} else {
				m[fmt.Sprintf("b%d-%d-%d", r/3, c/3, item)] = struct{}{}
			}
		}
	}
	return true
}
