package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		desc string
		m    int
		n    int
		want int
	}{
		{
			desc: "#1",
			m:    1,
			n:    1,
			want: 1,
		},
		{
			desc: "#2",
			m:    2,
			n:    1,
			want: 1,
		},
		{
			desc: "#3",
			m:    1,
			n:    2,
			want: 1,
		},
		{
			desc: "#4",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			desc: "#5",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			desc: "#6",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			desc: "#7",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			desc: "#8",
			m:    51,
			n:    9,
			want: 1916797311,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := uniquePaths2(s.m, s.n); got != s.want {
				t.Errorf("m: %d, n: %d, want: %d, got: %d", s.m, s.n, s.want, got)
			}
		})
	}
}
