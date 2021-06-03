package leetcode

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want [][]int
	}{
		{
			desc: "#1",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {2, 3}, {1, 2}, {1, 3}, {1, 2, 3}},
		},
		{
			desc: "#2",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := subsets(s.nums); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
