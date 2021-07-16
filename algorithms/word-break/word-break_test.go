package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	cases := []struct {
		desc     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			desc:     "#1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			desc:     "#2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			desc:     "#3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			desc:     "#4",
			s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			want:     false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := wordBreak(s.s, s.wordDict); got != s.want {
				t.Errorf("s: %s, wordDict: %v, want: %v, got: %v", s.s, s.wordDict, s.want, got)
			}
		})
	}
}
