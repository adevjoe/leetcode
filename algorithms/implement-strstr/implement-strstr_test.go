package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	cases := []struct {
		desc     string
		haystack string
		needle   string
		want     int
	}{
		{
			desc:     "#1",
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			desc:     "#2",
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			desc:     "#3",
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			desc:     "#4",
			haystack: "",
			needle:   "a",
			want:     -1,
		},
		{
			desc:     "#5",
			haystack: "aaa",
			needle:   "aaa",
			want:     0,
		},
		{
			desc:     "#5",
			haystack: "a",
			needle:   "",
			want:     0,
		},
		{
			desc:     "#6",
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
		{
			desc:     "#7",
			haystack: "mississippi",
			needle:   "issi",
			want:     1,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := strStr(s.haystack, s.needle); got != s.want {
				t.Errorf("haystack: %s, needle: %s, want: %d, got: %d", s.haystack, s.needle, s.want, got)
			}
		})
	}
}
