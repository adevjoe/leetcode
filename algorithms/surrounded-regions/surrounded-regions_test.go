package leetcode

import (
	"reflect"
	"testing"
)

func TestSurroundedRegions(t *testing.T) {
	cases := []struct {
		desc  string
		board [][]byte
		want  [][]byte
	}{
		{
			desc:  "#1",
			board: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			want:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'X', 'X', 'X'}},
		},
		{
			desc:  "#2",
			board: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X'}},
			want:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X'}},
		},
		{
			desc:  "#3",
			board: [][]byte{{'O', 'X', 'O'}, {'X', 'O', 'X'}, {'O', 'X', 'O'}},
			want:  [][]byte{{'O', 'X', 'O'}, {'X', 'X', 'X'}, {'O', 'X', 'O'}},
		},
		{
			desc:  "#4",
			board: [][]byte{{'X'}},
			want:  [][]byte{{'X'}},
		},
		{
			desc:  "#5",
			board: [][]byte{{'O'}},
			want:  [][]byte{{'O'}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			origin := make([][]byte, len(s.board))
			copy(origin, s.board)
			if solve(s.board); !reflect.DeepEqual(s.board, origin) {
				t.Error(origin)
			}
		})
	}
}
