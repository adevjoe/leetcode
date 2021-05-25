package leetcode

import (
	"reflect"
	"testing"
)

func TestSetMatrixZeroes(t *testing.T) {
	cases := []struct {
		desc   string
		matrix [][]int
		want   [][]int
	}{
		{
			desc: "#1",
			matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			desc: "#2",
			matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			desc: "#3",
			matrix: [][]int{
				{0},
			},
			want: [][]int{
				{0},
			},
		},
		{
			desc: "#4",
			matrix: [][]int{
				{1},
			},
			want: [][]int{
				{1},
			},
		},
		{
			desc: "#5",
			matrix: [][]int{
				{1, 0, 1},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			desc: "#6",
			matrix: [][]int{
				{1},
				{1},
				{0},
			},
			want: [][]int{
				{0},
				{0},
				{0},
			},
		},
		{
			desc: "#7",
			matrix: [][]int{
				{1, 1, 1},
				{0, 1, 1},
				{0, 2, 3},
			},
			want: [][]int{
				{0, 1, 1},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	cases2 := make([]struct {
		desc   string
		matrix [][]int
		want   [][]int
	}, 0)
	copy(cases2, cases)

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if setZeroes(s.matrix); !reflect.DeepEqual(s.matrix, s.want) {
				t.Error()
			}
		})
	}
	for _, s := range cases2 {
		t.Run(s.desc, func(t *testing.T) {
			if setZeroes2(s.matrix); !reflect.DeepEqual(s.matrix, s.want) {
				t.Error()
			}
		})
	}
}
