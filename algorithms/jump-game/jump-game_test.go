package leetcode

import "testing"

func TestJumpGame(t *testing.T) {
	cases := []struct {
		desc string
		nums []int
		want bool
	}{
		{
			desc: "#1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			desc: "#2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			desc: "#3",
			nums: []int{1, 2, 0, 1},
			want: true,
		},
		{
			desc: "#4",
			nums: []int{2, 0, 0},
			want: true,
		},
		{
			desc: "#5",
			nums: []int{0},
			want: true,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := canJump(s.nums); got != s.want {
				t.Errorf("nums: %v, want: %v, got: %v", s.nums, s.want, got)
			}
			if got := canJump2(s.nums); got != s.want {
				t.Errorf("nums: %v, want: %v, got: %v", s.nums, s.want, got)
			}
		})
	}
}
