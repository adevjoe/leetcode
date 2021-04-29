package leetcode

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	cases := []struct {
		desc   string
		matrix [][]int
		want   [][]int
	}{
		{
			desc: "#1",
			matrix: [][]int{
				{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1}, {8, 5, 2}, {9, 6, 3},
			},
		},
		{
			desc: "#2",
			matrix: [][]int{
				{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
			},
			want: [][]int{
				{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11},
			},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			var matrix [][]int
			copy(matrix, s.matrix)
			if rotate(s.matrix); !reflect.DeepEqual(s.matrix, s.want) {
				t.Errorf("matrix: %v, want: %v, got: %v", matrix, s.want, s.matrix)
			}
		})
	}
}
