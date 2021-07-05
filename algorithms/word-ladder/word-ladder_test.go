package leetcode

import "testing"

func TestWordLadder(t *testing.T) {
	cases := []struct {
		desc      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			desc:      "#1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			desc:      "#2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := ladderLength(s.beginWord, s.endWord, s.wordList); got != s.want {
				t.Error()
			}
		})
	}
}
