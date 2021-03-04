package leetcode

import "testing"

func TestLongestWordinDictionarythroughDeleting(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		d    []string
		want string
	}{
		{
			desc: "not found",
			s:    "abcd",
			d:    []string{"ef", "g"},
			want: "",
		},
		{
			desc: "found",
			s:    "abpcplea",
			d:    []string{"ale", "apple", "monkey", "plea"},
			want: "apple",
		},
		{
			desc: "found first one",
			s:    "abpcplea",
			d:    []string{"a", "b", "c"},
			want: "a",
		},
		{
			desc: "smallest lexicographical order",
			s:    "bab",
			d:    []string{"ba", "ab", "a", "b"},
			want: "ab",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := findLongestWord(s.s, s.d); got != s.want {
				t.Errorf("s: %s, d: %v, want: %s, got: %s", s.s, s.d, s.want, got)
			}
		})
	}
}
