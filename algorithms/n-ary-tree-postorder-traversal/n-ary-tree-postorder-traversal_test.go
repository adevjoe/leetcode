package leetcode

import (
	"github.com/adevjoe/leetcode/common"
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	cases := []struct {
		desc string
		root *common.Node
		want []int
	}{
		{
			desc: "#1",
			root: &common.Node{
				Val: 1,
				Children: []*common.Node{
					{Val: 3, Children: []*common.Node{{Val: 5}, {Val: 6}}},
					{Val: 2},
					{Val: 4},
				},
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
		{
			desc: "#2",
			root: &common.Node{
				Val: 1,
				Children: []*common.Node{
					{Val: 2},
					{Val: 3, Children: []*common.Node{{Val: 6}, {Val: 7, Children: []*common.Node{{Val: 11, Children: []*common.Node{{Val: 14}}}}}}},
					{Val: 4, Children: []*common.Node{{Val: 8, Children: []*common.Node{{Val: 12}}}}},
					{Val: 5, Children: []*common.Node{{Val: 9, Children: []*common.Node{{Val: 13}}}, {Val: 10}}},
				},
			},
			want: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := postorder(s.root); !reflect.DeepEqual(got, s.want) {
				t.Errorf("want: %v, got: %v", s.want, got)
			}
		})
	}
}
