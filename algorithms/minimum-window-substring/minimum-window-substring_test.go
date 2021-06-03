package leetcode

import "testing"

func TestMinimumWindowSubstring(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		t    string
		want string
	}{
		{
			desc: "#1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			desc: "#2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			desc: "#3",
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := minWindow(s.s, s.t); s.want != got {
				t.Errorf("s: %s, t: %s, want: %s, got: %s", s.s, s.t, s.want, got)
			}
		})
	}
}
