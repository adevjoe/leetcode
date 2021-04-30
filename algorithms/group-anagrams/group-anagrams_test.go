package leetcode

import (
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		desc string
		strs []string
		want [][]string
	}{
		{
			desc: "#1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			desc: "#2",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			desc: "#3",
			strs: []string{"a"},
			want: [][]string{{"a"}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := groupAnagrams(s.strs); len(got) != len(s.want) {
				t.Error()
			}
		})
	}
}
