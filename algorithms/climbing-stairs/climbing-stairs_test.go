package leetcode

import "testing"

func TestClimbingStairs(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		want int
	}{
		{
			desc: "#1",
			n:    1,
			want: 1,
		},
		{
			desc: "#2",
			n:    2,
			want: 2,
		},
		{
			desc: "#3",
			n:    3,
			want: 3,
		},
		{
			desc: "#4",
			n:    20,
			want: 10946,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := climbStairs(s.n); got != s.want {
				t.Errorf("n: %d, want: %d, got: %d", s.n, s.want, got)
			}
			if got := climbStairs2(s.n); got != s.want {
				t.Errorf("n: %d, want: %d, got: %d", s.n, s.want, got)
			}
		})
	}
}
