// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// 下面两种写法效率差不多，在 leetcode 找到的其它解法也是类似，但是 Runtime 只是比 85% 的提交快，感觉应该是以前的用例少一点。优化空间已经很小了。

package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if (p.Val <= root.Val && root.Val <= q.Val) || (p.Val >= root.Val && root.Val >= q.Val) {
		return root
	}
	if root.Left != nil {
		l := lowestCommonAncestor(root.Left, p, q)
		if l != nil {
			return l
		}
	}
	if root.Right != nil {
		r := lowestCommonAncestor(root.Right, p, q)
		if r != nil {
			return r
		}
	}
	return nil
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if (p.Val <= root.Val && root.Val <= q.Val) || (p.Val >= root.Val && root.Val >= q.Val) {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		if root.Left != nil {
			l := lowestCommonAncestor(root.Left, p, q)
			if l != nil {
				return l
			}
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		if root.Right != nil {
			r := lowestCommonAncestor(root.Right, p, q)
			if r != nil {
				return r
			}
		}
	}
	return nil
}
