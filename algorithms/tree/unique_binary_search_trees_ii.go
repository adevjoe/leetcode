// https://leetcode.com/problems/unique-binary-search-trees-ii/

package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
n = 1
1

n = 2
  1         2
    2     1

n = 3

  1         1          2            3         3
     2         3     1   3       2         1
        3    2                 1             2
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return gen(1, n)
}

func gen(start, end int) []*TreeNode {
	var trees []*TreeNode
	if start > end {
		trees = append(trees, nil)
		return trees
	}

	if start == end {
		trees = append(trees, &TreeNode{
			Val:   start,
			Left:  nil,
			Right: nil,
		})
		return trees
	}

	for i := start; i <= end; i++ {
		left := gen(start, i-1)
		right := gen(i+1, end)

		for _, lNode := range left {
			for _, rNode := range right {
				root := &TreeNode{
					Val:   i,
					Left:  lNode,
					Right: rNode,
				}
				trees = append(trees, root)
			}
		}
	}
	return trees
}
