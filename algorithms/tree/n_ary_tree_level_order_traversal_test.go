package tree

import (
	"reflect"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	cases := []struct {
		desc string
		root *NodeOne
		want [][]int
	}{
		{
			desc: "#1",
			root: &NodeOne{Val: 1, Children: []*NodeOne{{Val: 3, Children: []*NodeOne{{Val: 5}, {Val: 6}}}, {Val: 2}, {Val: 4}}},
			want: [][]int{{1}, {3, 2, 4}, {5, 6}},
		},
		{
			desc: "#2",
			root: &NodeOne{
				Val: 1,
				Children: []*NodeOne{
					{Val: 2},
					{Val: 3, Children: []*NodeOne{
						{Val: 6},
						{Val: 7, Children: []*NodeOne{{Val: 11, Children: []*NodeOne{{Val: 14}}}}},
					}},
					{Val: 4, Children: []*NodeOne{{Val: 8, Children: []*NodeOne{{Val: 12}}}}},
					{Val: 5, Children: []*NodeOne{{Val: 9, Children: []*NodeOne{{Val: 13}}}, {Val: 10}}},
				},
			},
			want: [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := levelOrder(s.root); !reflect.DeepEqual(got, s.want) {
				t.Errorf("got: %v, want, %v", got, s.want)
			}
		})
	}
}
