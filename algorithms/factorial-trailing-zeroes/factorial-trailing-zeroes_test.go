package leetcode

import "testing"

func TestFactorialTrailingZeroes(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		want int
	}{
		{
			desc: "#1",
			n:    0,
			want: 0,
		},
		{
			desc: "#2",
			n:    1,
			want: 0,
		},
		{
			desc: "#3",
			n:    5,
			want: 1,
		},
		{
			desc: "#4",
			n:    25,
			want: 6,
		},
		{
			desc: "#5",
			n:    300,
			want: 74,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := trailingZeroes(s.n); got != s.want {
				t.Error()
			}
		})
	}
}
