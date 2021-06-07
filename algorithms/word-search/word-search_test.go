package leetcode

import "testing"

func TestWordSearch(t *testing.T) {
	cases := []struct {
		desc  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			desc: "#1",
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word: "ABCCED",
			want: true,
		},
		{
			desc: "#2",
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word: "SEE",
			want: true,
		},
		{
			desc: "#3",
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word: "ABCB",
			want: false,
		},
		{
			desc: "#4",
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SCCS"),
				[]byte("ADEE"),
			},
			word: "ABCCED",
			want: true,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := exist(s.board, s.word); got != s.want {
				t.Error()
			}
		})
	}
}
