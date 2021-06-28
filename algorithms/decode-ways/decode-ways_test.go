package leetcode

import "testing"

func TestDecodeWays(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		want int
	}{
		{
			desc: "#1",
			s:    "12",
			want: 2,
		},
		{
			desc: "#2",
			s:    "226",
			want: 3,
		},
		{
			desc: "#3",
			s:    "0",
			want: 0,
		},
		{
			desc: "#4",
			s:    "06",
			want: 0,
		},
		{
			desc: "#5",
			s:    "10",
			want: 1,
		},
		{
			desc: "#6",
			s:    "22612",
			want: 6,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := numDecodings(s.s); got != s.want {
				t.Error()
			}
		})
	}
}
