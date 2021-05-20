package leetcode

import "testing"

func TestSqrt(t *testing.T) {
	cases := []struct {
		desc string
		x    int
		want int
	}{
		{
			desc: "#1",
			x:    0,
			want: 0,
		},
		{
			desc: "#2",
			x:    1,
			want: 1,
		},
		{
			desc: "#3",
			x:    2,
			want: 1,
		},
		{
			desc: "#4",
			x:    4,
			want: 2,
		},
		{
			desc: "#5",
			x:    8,
			want: 2,
		},
		{
			desc: "#6",
			x:    523423423,
			want: 22878,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := mySqrt(s.x); got != s.want {
				t.Errorf("x: %d, want: %d, got: %d", s.x, s.want, got)
			}
		})
	}
}
