package leetcode

import "testing"

func TestHappyNumber(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		want bool
	}{
		{
			desc: "#1",
			n:    19,
			want: true,
		},
		{
			desc: "#2",
			n:    2,
			want: false,
		},
		{
			desc: "#3",
			n:    7,
			want: true,
		},
		{
			desc: "#4",
			n:    16,
			want: false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isHappy(s.n); got != s.want {
				t.Error()
			}
			if got := isHappy2(s.n); got != s.want {
				t.Error()
			}
		})
	}
}
