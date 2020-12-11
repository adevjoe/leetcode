// https://leetcode.com/problems/n-ary-tree-postorder-traversal/

package algorithms

import "github.com/adevjoe/leetcode/common"

func postorder(root *common.Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	for _, r := range root.Children {
		result = append(result, postorder(r)...)
	}
	result = append(result, root.Val)
	return result
}
