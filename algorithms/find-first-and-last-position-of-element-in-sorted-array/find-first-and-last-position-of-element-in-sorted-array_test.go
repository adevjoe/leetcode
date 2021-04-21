package leetcode

import (
	"reflect"
	"testing"
)

func TestFindFirstandLastPositionofElementinSortedArray(t *testing.T) {
	cases := []struct {
		desc   string
		nums   []int
		target int
		want   []int
	}{
		{
			desc:   "#1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			desc:   "one",
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			desc:   "two",
			nums:   []int{2, 2},
			target: 2,
			want:   []int{0, 1},
		},
		{
			desc:   "not found",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			desc:   "nil",
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := searchRange(s.nums, s.target); !reflect.DeepEqual(s.want, got) {
				t.Errorf("nums: %v, target: %d, want: %v, got: %v", s.nums, s.target, s.want, got)
			}
			if got := searchRange1(s.nums, s.target); !reflect.DeepEqual(s.want, got) {
				t.Errorf("nums: %v, target: %d, want: %v, got: %v", s.nums, s.target, s.want, got)
			}
		})
	}
}
