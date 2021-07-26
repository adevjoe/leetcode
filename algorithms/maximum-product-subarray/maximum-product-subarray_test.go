package leetcode

import "testing"

func TestMaximumProductSubarray(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			desc: "#2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := maxProduct(s.nums); got != s.want {
				t.Error()
			}
		})
	}
}
