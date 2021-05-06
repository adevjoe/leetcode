package leetcode

import "testing"

func TestPow(t *testing.T) {
	cases := []struct {
		desc string
		x    float64
		n    int
		want float64
	}{
		{
			desc: "#1",
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			desc: "#2",
			x:    2.1,
			n:    3,
			want: 9.261,
		},
		{
			desc: "#3",
			x:    2,
			n:    -2,
			want: 0.25,
		},
		{
			desc: "#4",
			x:    8.88023,
			n:    3,
			want: 700.28148,
		},
		{
			desc: "#5",
			x:    0.44894,
			n:    -5,
			want: 54.83508,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := myPow(s.x, s.n); got != s.want {
				t.Errorf("x: %v, n: %d, want: %v, got: %v", s.x, s.n, s.want, got)
			}
		})
	}
}
