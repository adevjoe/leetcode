package leetcode

import (
	"reflect"
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	cases := []struct {
		desc    string
		numRows int
		want    [][]int
	}{
		{
			desc:    "#1",
			numRows: 1,
			want:    [][]int{{1}},
		},
		{
			desc:    "#2",
			numRows: 2,
			want:    [][]int{{1}, {1, 1}},
		},
		{
			desc:    "#3",
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := generate(s.numRows); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
}
