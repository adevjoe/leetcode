package leetcode

import (
	"reflect"
	"testing"
)

func TestPalindromePartitioning(t *testing.T) {
	cases := []struct {
		desc string
		s    string
		want [][]string
	}{
		{
			desc: "#1",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			desc: "#2",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			desc: "#3",
			s:    "aa",
			want: [][]string{{"a", "a"}, {"aa"}},
		},
		{
			desc: "#4",
			s:    "ab",
			want: [][]string{{"a", "b"}},
		},
		{
			desc: "#5",
			s:    "accadeffe",
			want: [][]string{
				{"a", "c", "c", "a", "d", "e", "f", "f", "e"}, {"a", "c", "c", "a", "d", "e", "ff", "e"}, {"a", "c", "c", "a", "d", "effe"}, {"a", "cc", "a", "d", "e", "f", "f", "e"}, {"a", "cc", "a", "d", "e", "ff", "e"}, {"a", "cc", "a", "d", "effe"}, {"acca", "d", "e", "f", "f", "e"}, {"acca", "d", "e", "ff", "e"}, {"acca", "d", "effe"}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := partition(s.s); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
}
