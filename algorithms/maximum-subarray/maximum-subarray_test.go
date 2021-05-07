package leetcode

import "testing"

func TestMaximumSubarray(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		{
			desc: "#2",
			nums: []int{1},
			want: 1,
		},
		{
			desc: "#3",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			desc: "#4",
			nums: []int{5, 4, -10, 7, 8},
			want: 15,
		},
		{
			desc: "#5",
			nums: []int{-1},
			want: -1,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := maxSubArray(s.nums); got != s.want {
				t.Errorf("nums: %v, want: %d, got: %d", s.nums, s.want, got)
			}
		})
	}
}
