// https://leetcode.com/problems/house-robber-iii/

package house_robber

import "math"

var robThreeMap map[*TreeNode]map[bool]int

func robThree(root *TreeNode) int {
	return robThreeSub(root, true)
}

func robThreeTwo(root *TreeNode) int {
	robThreeMap = make(map[*TreeNode]map[bool]int)
	return robThreeSubTwo(root, true)
}

func robThreeSub(root *TreeNode, haveRoot bool) int {
	if root == nil {
		return 0
	}
	if haveRoot {
		a := root.Val + robThreeSub(root.Left, false) + robThreeSub(root.Right, false)
		b := robThreeSub(root.Left, true) + robThreeSub(root.Right, true)
		return int(math.Max(float64(a), float64(b)))
	} else {
		return robThreeSub(root.Left, true) + robThreeSub(root.Right, true)
	}
}

func robThreeSubTwo(root *TreeNode, haveRoot bool) int {
	if root == nil {
		return 0
	}
	if robThreeMap[root][haveRoot] > 0 {
		return robThreeMap[root][haveRoot]
	}
	var max int
	if haveRoot {
		a := root.Val + robThreeSubTwo(root.Left, false) + robThreeSubTwo(root.Right, false)
		b := robThreeSubTwo(root.Left, true) + robThreeSubTwo(root.Right, true)
		max = int(math.Max(float64(a), float64(b)))
	} else {
		max = robThreeSubTwo(root.Left, true) + robThreeSubTwo(root.Right, true)
	}
	if robThreeMap[root] == nil {
		robThreeMap[root] = make(map[bool]int)
	}
	robThreeMap[root][haveRoot] = max
	return max
}
