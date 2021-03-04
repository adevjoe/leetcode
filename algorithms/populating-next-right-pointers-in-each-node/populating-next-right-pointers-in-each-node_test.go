package leetcode

import (
	"fmt"
	"testing"
)

func TestPopulatingNextRightPointersinEachNode(t *testing.T) {
	cases := []struct {
		arg1 *Node
		want *Node
	}{
		{arg1: &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Left: &Node{Val: 6}, Right: &Node{Val: 7}}}},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			connect(c.arg1)
		})
	}
}
