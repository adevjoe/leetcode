// https://leetcode.com/problems/generate-parentheses/

package generate_arentheses

import "github.com/adevjoe/leetcode/common"

/**
可以转为树结构来实现，树的深度为 2^n，根节点为 "("，然后往下为完全二叉树。
叶节点可以优化，因为最后一个括号只能为右括号，所以把所有为 "(" 的叶节点剪掉。
n = 1
// tree
  (
    )

// result
()

n = 2
// tree
     (
 (      )
( )    ( )
) )    ) )

// result
(()) ()()
*/

func generateParenthesis(n int) []string {
	var list []string
	for i := 1; i <= 2*n; i++ {
		last := false
		if i == 2*n {
			last = true
		}
		list = generate(list, last)
	}
	var result []string
	for _, s := range list {
		if isValid(s) {
			result = append(result, s)
		}
	}
	return result
}

func isValid(s string) bool {
	stack := common.NewStack()
	for _, v := range s {
		if v == '(' {
			stack.Push(')')
		} else if stack.Len() == 0 || stack.Pop() != v {
			return false
		}
	}
	return stack.Len() == 0
}

func generate(list []string, last bool) (result []string) {
	if len(list) == 0 {
		return []string{"("}
	}
	for _, s := range list {
		l := s + "("
		r := s + ")"
		if !last {
			result = append(result, l)
		}
		result = append(result, r)
	}
	return
}

// 在生成字符串过程中判断括号闭合
func generateParenthesis2(n int) []string {
	var result []string
	result = generate2(result, "", 0, 0, n)
	return result
}

func generate2(list []string, s string, open, closeP, max int) []string {
	if len(s) == 2*max {
		list = append(list, s)
		return list
	}
	if open < max {
		list = generate2(list, s+"(", open+1, closeP, max)
	}
	if closeP < open {
		list = generate2(list, s+")", open, closeP+1, max)
	}
	return list
}
