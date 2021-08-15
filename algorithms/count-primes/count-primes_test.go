package leetcode

import "testing"

func TestCountPrimes(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		want int
	}{
		{
			desc: "#1",
			n:    10,
			want: 4,
		},
		{
			desc: "#2",
			n:    0,
			want: 0,
		},
		{
			desc: "#3",
			n:    1,
			want: 0,
		},
		{
			desc: "#4",
			n:    2,
			want: 0,
		},
		{
			desc: "#5",
			n:    100,
			want: 25,
		},
		{
			desc: "#6",
			n:    200,
			want: 46,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := countPrimes(s.n); got != s.want {
				t.Error()
			}
		})
	}
}
