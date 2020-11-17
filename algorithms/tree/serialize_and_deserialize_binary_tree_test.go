package tree

import (
	"testing"
)

func TestSerializeAndDeserialize(t *testing.T) {
	cases := []struct {
		desc string
		arg  *TreeNode
	}{
		{
			desc: "#1",
			arg:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}},
		},
		{
			desc: "#2",
			arg:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}},
		},
		{
			desc: "null",
			arg:  nil,
		},
		{
			desc: "#4",
			arg:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		},
		{
			desc: "#5",
			arg:  &TreeNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			ser := ConstructorSD()
			deser := ConstructorSD()
			data := ser.serialize(s.arg)
			ans := deser.deserialize(data)
			if !isSameTree(s.arg, ans) {
				t.Errorf("arg: %v, string: %s, got: %v", inorderTraversal(s.arg), data, inorderTraversal(ans))
			}
		})
	}
}
