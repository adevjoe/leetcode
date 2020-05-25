// https://leetcode.com/problems/path-sum/

package tree

func pathSum(root *TreeNode, sum int) [][]int {
	var r [][]int
	if root == nil {
		return r
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	left := pathSum(root.Left, sum-root.Val)
	for i := range left {
		temp := left[i]
		left[i] = make([]int, 0)
		left[i] = append(left[i], root.Val)
		left[i] = append(left[i], temp...)
		r = append(r, left[i])
	}
	right := pathSum(root.Right, sum-root.Val)
	for i := range right {
		temp := right[i]
		right[i] = make([]int, 0)
		right[i] = append(right[i], root.Val)
		right[i] = append(right[i], temp...)
		r = append(r, right[i])
	}
	return r
}
