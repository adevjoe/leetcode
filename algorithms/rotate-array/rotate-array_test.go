package leetcode

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		k    int
		want []int
	}{
		{
			desc: "#1",
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			desc: "#2",
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			var origin []int
			copy(origin, s.nums)
			if rotate(s.nums, s.k); !reflect.DeepEqual(s.want, s.nums) {
				t.Error()
			}
		})
	}
}
