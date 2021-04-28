package leetcode

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want [][]int
	}{
		{
			desc: "#1",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			desc: "#2",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			desc: "#3",
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := permute(s.nums); !reflect.DeepEqual(got, s.want) {
				t.Errorf("nums: %v, want: %v, got: %v", s.nums, s.want, got)
			}
		})
	}
}
