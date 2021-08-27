package leetcode

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	cases := []struct {
		desc string
		arr  []int
		want []int
	}{
		{
			desc: "#1",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			desc: "#2",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			desc: "#3",
			arr:  []int{1, 0, 0, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 0, 0, 3, 0, 0},
		},
		{
			desc: "#4",
			arr:  []int{0},
			want: []int{0},
		},
		{
			desc: "#5",
			arr:  []int{1},
			want: []int{1},
		},
		{
			desc: "#6",
			arr:  []int{0, 0, 0, 0, 0, 0, 0},
			want: []int{0, 0, 0, 0, 0, 0, 0},
		},
		{
			desc: "#7",
			arr:  []int{8, 4, 5, 0, 0, 0, 0, 7},
			want: []int{8, 4, 5, 0, 0, 0, 0, 0},
		},
		{
			desc: "#8",
			arr:  []int{0, 1, 7, 6, 0, 2, 0, 7},
			want: []int{0, 0, 1, 7, 6, 0, 0, 2},
		},
		{
			desc: "#9",
			arr:  []int{9, 9, 9, 4, 8, 0, 0, 3, 7, 2, 0, 0, 0, 0, 9, 1, 0, 0, 1, 1, 0, 5, 6, 3, 1, 6, 0, 0, 2, 3, 4, 7, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 7, 3, 3, 0, 5, 7, 0, 8, 1, 9, 6, 3, 0, 8, 8, 8, 8, 0, 0, 5, 0, 0, 0, 3, 7, 7, 7, 7, 5, 1, 0, 0, 8, 0, 0},
			want: []int{9, 9, 9, 4, 8, 0, 0, 0, 0, 3, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 1, 1, 0, 0, 5, 6, 3, 1, 6, 0, 0, 0, 0, 2, 3, 4, 7, 0, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 0, 0, 7, 3, 3, 0, 0, 5, 7, 0, 0, 8, 1, 9, 6, 3, 0, 0, 8, 8, 8, 8, 0},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			arr1 := make([]int, len(s.arr))
			copy(arr1, s.arr)
			if duplicateZeros(arr1); !reflect.DeepEqual(s.want, arr1) {
				t.Errorf("arr: %v, got: %v, want: %v", s.arr, arr1, s.want)
			}
			arr2 := make([]int, len(s.arr))
			copy(arr2, s.arr)
			if duplicateZeros2(arr2); !reflect.DeepEqual(s.want, arr2) {
				t.Errorf("arr: %v, got: %v, want: %v", s.arr, arr2, s.want)
			}
		})
	}
}
