package leetcode

import "testing"

func TestWildcardMatching(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		p    string
		want bool
	}{
		{
			desc: "#1",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			desc: "#2",
			s:    "aa",
			p:    "*",
			want: true,
		},
		{
			desc: "#3",
			s:    "cb",
			p:    "?a",
			want: false,
		},
		{
			desc: "#4",
			s:    "adceb",
			p:    "*a*b",
			want: true,
		},
		{
			desc: "#5",
			s:    "acdcb",
			p:    "a*c?b",
			want: false,
		},
		{
			desc: "#6",
			s:    "",
			p:    "*",
			want: true,
		},
		{
			desc: "#7",
			s:    "",
			p:    "?*",
			want: false,
		},
		{
			desc: "#8",
			s:    "abccaadbc",
			p:    "*a*b*",
			want: true,
		},
		{
			desc: "#9",
			s:    "abccaadbc",
			p:    "*ab*",
			want: true,
		},
		{
			desc: "#10",
			s:    "abccaadbc",
			p:    "*ad*",
			want: true,
		},
		{
			desc: "#11",
			s:    "abccaadbc",
			p:    "*a*d",
			want: false,
		},
		{
			desc: "#12",
			s:    "a",
			p:    "",
			want: false,
		},
		{
			desc: "#13",
			s:    "abcdef",
			p:    "a?de*",
			want: false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isMatch(s.s, s.p); got != s.want {
				t.Errorf("s: %s, p: %s, want: %v, got: %v", s.s, s.p, s.want, got)
			}
		})
	}
}
