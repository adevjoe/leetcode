package leetcode

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want int
	}{
		{
			desc: "#1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			desc: "#2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			desc: "#3",
			nums: []int{-6, 8, -5, 7, -9, -1, -7, -6, -9, -7, 5, 7, -1, -8, -8, -2, 0},
			want: 5,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := longestConsecutive(s.nums); got != s.want {
				t.Error()
			}
			if got := longestConsecutive2(s.nums); got != s.want {
				t.Error()
			}
		})
	}
}
