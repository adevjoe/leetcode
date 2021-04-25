package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			desc: "#2",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			desc: "#3",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			desc: "#4",
			nums: []int{1, 1},
			want: 2,
		},
		{
			desc: "#5",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			desc: "#6",
			nums: []int{-1, 4, 2, 1, 9, 10},
			want: 3,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := firstMissingPositive(s.nums); got != s.want {
				t.Errorf("nums: %v, want: %d, got: %d", s.nums, s.want, got)
			}
			if got := firstMissingPositive2(s.nums); got != s.want {
				t.Errorf("nums: %v, want: %d, got: %d", s.nums, s.want, got)
			}
		})
	}
}
