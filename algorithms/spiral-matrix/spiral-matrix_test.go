package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	cases := []struct {
		desc   string
		matrix [][]int
		want   []int
	}{
		{
			desc:   "#1",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			desc:   "#2",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want:   []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			desc:   "#3",
			matrix: [][]int{{1}},
			want:   []int{1},
		},
		{
			desc:   "#4",
			matrix: [][]int{{1, 2, 3}},
			want:   []int{1, 2, 3},
		},
		{
			desc:   "#5",
			matrix: [][]int{{1}, {2}},
			want:   []int{1, 2},
		},
		{
			desc:   "#6",
			matrix: [][]int{{1, 2}, {3, 4}},
			want:   []int{1, 2, 4, 3},
		},
		{
			desc:   "#7",
			matrix: [][]int{{7}, {9}, {6}},
			want:   []int{7, 9, 6},
		},
		{
			desc:   "#8",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			want:   []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := spiralOrder(s.matrix); !reflect.DeepEqual(s.want, got) {
				t.Errorf("matrix: %v, want: %v, got: %v", s.matrix, s.want, got)
			}
		})
	}
}
