package leetcode

import "testing"

func TestValidPalindrome(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		want bool
	}{
		{
			desc: "#1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			desc: "#2",
			s:    "race a car",
			want: false,
		},
		{
			desc: "#3",
			s:    "a",
			want: true,
		},
		{
			desc: "#4",
			s:    ",",
			want: true,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isPalindrome(s.s); got != s.want {
				t.Error()
			}
		})
	}
}
